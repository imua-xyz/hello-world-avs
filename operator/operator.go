package operator

import (
	"context"
	"encoding/hex"
	"fmt"
	avs "github.com/ExocoreNetwork/exocore-avs/contracts/bindings/avs"
	"github.com/ExocoreNetwork/exocore-avs/core"
	chain "github.com/ExocoreNetwork/exocore-avs/core/chainio"
	"github.com/ExocoreNetwork/exocore-avs/core/chainio/eth"
	"github.com/ExocoreNetwork/exocore-avs/types"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	"github.com/ExocoreNetwork/exocore-sdk/crypto/bls"
	sdklogging "github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	blscommon "github.com/prysmaticlabs/prysm/v4/crypto/bls/common"
	"math/big"
	"time"

	use "github.com/ExocoreNetwork/exocore-avs/avs"
	"github.com/ExocoreNetwork/exocore-sdk/nodeapi"
	"github.com/ExocoreNetwork/exocore-sdk/signerv2"
	"github.com/ethereum/go-ethereum/common"
	"os"
)

const (
	AvsName = "hello-world-avs-demo"
	SemVer  = "0.0.1"
	// DayEpochID defines the identifier for a daily epoch.
	DayEpochID = "day"
	// HourEpochID defines the identifier for an hourly epoch.
	HourEpochID = "hour"
	// MinuteEpochID defines the identifier for an epoch that is a minute long.
	MinuteEpochID = "minute"
	// WeekEpochID defines the identifier for a weekly epoch.
	WeekEpochID = "week"
)

type Operator struct {
	config        types.NodeConfig
	logger        sdklogging.Logger
	ethClient     eth.EthClient
	nodeApi       *nodeapi.NodeApi
	avsWriter     chain.EXOWriter
	avsReader     chain.EXOChainReader
	avsSubscriber chain.AvsRegistrySubscriber

	blsKeypair   blscommon.SecretKey
	operatorAddr common.Address
	// receive new tasks in this chan (typically from listening to onchain event)
	newTaskCreatedChan chan *avs.ContracthelloWorldTaskCreated
	// needed when opting in to avs (allow this service manager contract to slash operator)
	avsAddr         common.Address
	epochIdentifier string
}

func NewOperatorFromConfig(c types.NodeConfig) (*Operator, error) {
	var logLevel sdklogging.LogLevel
	if c.Production {
		logLevel = sdklogging.Production
	} else {
		logLevel = sdklogging.Development
	}
	logger, err := sdklogging.NewZapLogger(logLevel)
	if err != nil {
		return nil, err
	}

	// Setup Node Api
	nodeApi := nodeapi.NewNodeApi(AvsName, SemVer, c.NodeApiIpPortAddress, logger)

	var ethRpcClient, ethWsClient eth.EthClient
	ethRpcClient, err = eth.NewClient(c.EthRpcUrl)
	if err != nil {
		logger.Error("can not create http eth client", "err", err)

		return nil, err
	}
	ethWsClient, err = eth.NewClient(c.EthWsUrl)
	if err != nil {
		logger.Error("Cannot create ws eth client", "err", err)
		return nil, err
	}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		logger.Info("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(c.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		logger.Error("Cannot parse bls private key", "err", err)
		return nil, err
	}

	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		logger.Info("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: c.OperatorEcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil {
		panic(err)
	}

	txMgr := txmgr.NewSimpleTxManager(ethRpcClient, logger, signerV2, common.HexToAddress(c.OperatorAddress))

	avsReader, _ := chain.BuildExoChainReader(
		common.HexToAddress(c.AVSAddress),
		ethRpcClient,
		logger)

	avsWriter, _ := chain.BuildExoChainWriter(
		common.HexToAddress(c.AVSAddress),
		ethRpcClient,
		logger,
		txMgr)

	avsSubscriber, _ := chain.BuildAvsRegistryChainSubscriber(
		common.HexToAddress(c.AVSAddress),
		ethWsClient,
		logger,
	)

	if err != nil {
		logger.Error("Cannot create AvsSubscriber", "err", err)
		return nil, err
	}
	epochIdentifier, err := avsReader.GetAVSInfo(&bind.CallOpts{}, c.AVSAddress)
	if err != nil {
		logger.Error("Cannot GetAVSInfo", "err", err)
		return nil, err
	}
	operator := &Operator{
		config:             c,
		logger:             logger,
		nodeApi:            nodeApi,
		ethClient:          ethRpcClient,
		avsWriter:          avsWriter,
		avsReader:          *avsReader,
		avsSubscriber:      avsSubscriber,
		blsKeypair:         blsKeyPair,
		operatorAddr:       common.HexToAddress(c.OperatorAddress),
		newTaskCreatedChan: make(chan *avs.ContracthelloWorldTaskCreated),
		avsAddr:            common.HexToAddress(c.AVSAddress),
		epochIdentifier:    epochIdentifier,
	}

	if c.RegisterOperatorOnStartup {

		if err != nil {
			return nil, err
		}
		operator.registerOperatorOnStartup()
	}

	logger.Info("Operator info",
		"operatorAddr", c.OperatorAddress,
		"operatorKey", operator.blsKeypair.PublicKey().Marshal(),
	)

	return operator, nil

}

func (o *Operator) Start(ctx context.Context) error {
	// 1.operator register  exocore
	// 2.operator optin avs
	// 3.operator accept staker delegation so that avs voting power is not 0, otherwise the task cannot be created
	// 4.operator register BLSPublicKey
	// 5.operator sumit task parse

	flag, err := o.avsReader.IsOperator(&bind.CallOpts{}, o.operatorAddr.String())
	if err != nil {
		o.logger.Error("Cannot exec IsOperator", "err", err)
		return err
	}
	if !flag {
		o.logger.Info("Operator is not registered.")
		_, err = o.avsWriter.RegisterOperatorToExocore(context.Background(), use.GenerateRandomName(8))
		if err != nil {
			o.logger.Error("Avs failed to RegisterOperatorToExocore", "err", err)
			return err
		}
	}

	pubKey, err := o.avsReader.GetRegisteredPubkey(&bind.CallOpts{}, o.operatorAddr.String())
	if err != nil {
		o.logger.Error("Cannot exec GetRegisteredPubKey", "err", err)
		return err
	}

	if pubKey == nil {
		// operator register BLSPublicKey  via evm tx
		msgBytes := crypto.Keccak256Hash([]byte("hello-avs")).Bytes()
		sig := o.blsKeypair.Sign(msgBytes)
		_, err = o.avsWriter.RegisterBLSPublicKey(
			context.Background(),
			o.operatorAddr.String(),
			o.blsKeypair.PublicKey().Marshal(),
			sig.Marshal(),
			msgBytes)

		if err != nil {
			o.logger.Error("operator failed to registerBLSPublicKey", "err", err)
			return err
		}
	}

	// check operator delegation usd amount

	amount, err := o.avsReader.GetOperatorOptedUSDValue(&bind.CallOpts{}, o.avsAddr.String(), o.operatorAddr.String())
	if err != nil {
		o.logger.Error("Cannot exec IsOperator", "err", err)
		return err
	}

	if amount.IsZero() {
		o.logger.Error("amount is zero,please delegate amount to the current operator  ", "amount", amount)
		return err
	}
	o.logger.Infof("Starting operator.")

	if o.config.EnableNodeApi {
		o.nodeApi.Start()
	}

	firstHeight, err := o.ethClient.BlockNumber(context.Background())
	if err != nil {
		o.logger.Error("Cannot create AvsSubscriber", "err", err)
		return err
	}
	o.GetLog(int64(firstHeight))

	height := firstHeight
	o.logger.Info("Event firstHeight: %v\n", firstHeight)

	for {
		currentHeight, err := o.ethClient.BlockNumber(context.Background())
		o.logger.Info("Event currentHeight: %v\n", currentHeight)

		if err != nil {
			o.logger.Fatal(err.Error())
		}
		if currentHeight == height+1 {
			o.GetLog(int64(currentHeight))

			height = currentHeight
		}
		time.Sleep(2 * time.Second) // Wait for 2 seconds and check again

	}

}
func (o *Operator) GetLog(height int64) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{o.avsAddr},
		FromBlock: big.NewInt(height),
		ToBlock:   big.NewInt(height),
	}

	logs, err := o.ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		o.logger.Fatal(err.Error())
	}
	if logs != nil {
		contractAbi, _ := avs.ContracthelloWorldMetaData.GetAbi()
		event := contractAbi.Events["TaskCreated"]
		for _, vLog := range logs {
			data := vLog.Data

			eventArgs, err := event.Inputs.Unpack(data)
			if err != nil {
				o.logger.Fatal(err.Error())
			}
			if eventArgs != nil {
				taskResponse := o.ProcessNewTaskCreatedLog(eventArgs)
				sig, resBytes, err := o.SignTaskResponse(taskResponse)
				if err != nil {
					continue
				}
				o.logger.Info(string(sig))

				taskInfo, _ := o.avsReader.GetTaskInfo(&bind.CallOpts{}, o.avsAddr.String(), taskResponse.TaskID.Uint64())
				go o.SendSignedTaskResponseToExocore(taskResponse.TaskID.Uint64(), resBytes, sig, taskInfo)
			}
		}
	}

}

// ProcessNewTaskCreatedLog TaskResponse is the struct that is signed and sent to the exocore as a task response.
func (o *Operator) ProcessNewTaskCreatedLog(eventArgs []interface{}) *core.TaskResponse {
	o.logger.Debug("Received new task", "task", eventArgs)
	o.logger.Info("Received new task",
		"id", eventArgs[0].(*big.Int),
		"name", eventArgs[2].(string),
	)
	taskResponse := &core.TaskResponse{
		TaskID: eventArgs[0].(*big.Int),
		Msg:    eventArgs[2].(string),
	}
	return taskResponse
}

func (o *Operator) SignTaskResponse(taskResponse *core.TaskResponse) ([]byte, []byte, error) {

	taskResponseHash, err := core.GetTaskResponseDigestEncodeByjson(*taskResponse)
	if err != nil {
		o.logger.Error("Error SignTaskResponse with getting task response header hash. skipping task (this is not expected and should be investigated)", "err", err)
		return nil, nil, err
	}
	msgBytes := taskResponseHash[:]
	o.logger.Info("ResHash:", hex.EncodeToString(msgBytes))

	sig := o.blsKeypair.Sign(msgBytes)

	sigStr := hex.EncodeToString(sig.Marshal())
	o.logger.Info("sig:", sigStr)

	return sig.Marshal(), msgBytes, nil
}

func (o *Operator) SendSignedTaskResponseToExocore(
	taskId uint64,
	taskResponse []byte,
	blsSignature []byte,
	taskInfo []uint64) (string, error) {
	startingEpoch := taskInfo[0]
	taskResponsePeriod := taskInfo[1]
	taskStatisticalPeriod := taskInfo[2]

	for {
		epochIdentifier, err := o.avsReader.GetAVSInfo(&bind.CallOpts{}, o.avsAddr.String())
		if err != nil {
			o.logger.Error("Cannot GetAVSInfo", "err", err)
		}
		num, err := o.avsReader.GetCurrentEpoch(&bind.CallOpts{}, epochIdentifier)
		if err != nil {
			o.logger.Error("Cannot exec GetCurrentEpoch", "err", err)
		}
		currentEpoch := uint64(num)

		if currentEpoch > startingEpoch+taskResponsePeriod+taskStatisticalPeriod {
			break
		}
		o.logger.Info("Exit loop")
		switch {
		case currentEpoch <= startingEpoch:
			return "", fmt.Errorf("current epoch %d is less than starting epoch %d", currentEpoch, startingEpoch)
		case currentEpoch <= startingEpoch+taskResponsePeriod:
			tx, err := o.avsWriter.OperatorSubmitTask(
				context.Background(),
				taskId,
				nil,
				blsSignature,
				o.avsAddr.String(),
				"1")
			if err != nil {
				o.logger.Error("Avs failed to OperatorSubmitTask", "err", err)
			}
			return tx.BlockHash.String(), err
		case currentEpoch <= startingEpoch+taskStatisticalPeriod && currentEpoch > startingEpoch+taskResponsePeriod:
			tx, err := o.avsWriter.OperatorSubmitTask(
				context.Background(),
				taskId,
				taskResponse,
				blsSignature,
				o.avsAddr.String(),
				"2")
			if err != nil {
				o.logger.Error("Avs failed to OperatorSubmitTask", "err", err)
			}
			return tx.BlockHash.String(), err
		default:
			o.logger.Info("currentEpoch:", currentEpoch)
		}

		var sleepDuration time.Duration

		switch epochIdentifier {
		case DayEpochID:
			sleepDuration = 24 * time.Hour // Sleep for 24 hours (1 day)
		case HourEpochID:
			sleepDuration = time.Hour // Sleep for 1 hour
		case MinuteEpochID:
			sleepDuration = time.Minute // Sleep for 1 minute
		case WeekEpochID:
			sleepDuration = 7 * 24 * time.Hour // Sleep for 7 days (1 week)
		default:
			// Handle unknown epochIdentifier
			sleepDuration = 0
		}
		time.Sleep(sleepDuration)

	}
	return "", nil
}
