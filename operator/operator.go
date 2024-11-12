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
	"github.com/ExocoreNetwork/exocore-sdk/nodeapi"
	"github.com/ExocoreNetwork/exocore-sdk/signerv2"
	"github.com/cosmos/btcutil/bech32"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	blscommon "github.com/prysmaticlabs/prysm/v4/crypto/bls/common"
	"math/big"
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
	avsWriter     chain.ExoWriter
	avsReader     chain.ExoChainReader
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
	epochIdentifier, err := avsReader.GetAVSEpochIdentifier(&bind.CallOpts{}, c.AVSAddress)
	if err != nil {
		logger.Error("Cannot GetAVSEpochIdentifier", "err", err)
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
	// 1.operator register exocore
	// 2.operator opt-in avs
	// 3.operator accept staker delegation so that avs voting power is not 0, otherwise the task cannot be created
	// 4.operator register BLSPublicKey
	// 5.operator submit task

	operatorAddress, err := SwitchEthAddressToExoAddress(o.operatorAddr.String())
	if err != nil {
		o.logger.Error("Cannot switch eth address to exo address", "err", err)
		panic(err)
	}

	flag, err := o.avsReader.IsOperator(&bind.CallOpts{}, o.operatorAddr.String())
	if err != nil {
		o.logger.Error("Cannot exec IsOperator", "err", err)
		return err
	}
	if !flag {
		o.logger.Error("Operator is not registered.", "err", err)
		panic(fmt.Sprintf("Operator is not registered: %s", operatorAddress))
	}

	pubKey, err := o.avsReader.GetRegisteredPubkey(&bind.CallOpts{}, operatorAddress)
	if err != nil {
		o.logger.Error("Cannot exec GetRegisteredPubKey", "err", err)
		return err
	}

	if len(pubKey) == 0 {
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
		// return err
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
	// Channel to receive new block heights
	blockCh := make(chan uint64)

	// Start a goroutine to monitor block heights
	go func() {
		for {
			select {
			case <-ctx.Done():
				o.logger.Info("Stopping block monitoring")
				close(blockCh)
				return
			default:
				currentHeight, err := o.ethClient.BlockNumber(ctx)
				if err != nil {
					o.logger.Error("Error getting block number", "err", err)
					continue
				}
				if currentHeight == height+1 {
					blockCh <- currentHeight
					height = currentHeight
				}
			}
		}
	}()
	// Process new block heights as they are received
	for {
		select {
		case <-ctx.Done():
			o.logger.Info("Context cancelled, exiting loop")
			return nil
		case newHeight, ok := <-blockCh:
			if !ok {
				o.logger.Info("Block channel closed, exiting loop")
				return nil
			}
			if err := o.GetLog(int64(newHeight)); err != nil {
				o.logger.Error("Error retrieving logs for block %d: %v", newHeight, err)
			}
		}
	}
}
func (o *Operator) GetLog(height int64) error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{o.avsAddr},
		FromBlock: big.NewInt(height),
		ToBlock:   big.NewInt(height),
	}

	logs, err := o.ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		o.logger.Error("Failed to filter logs", "err", err)
		return err
	}
	if logs != nil {
		contractAbi, _ := avs.ContracthelloWorldMetaData.GetAbi()
		event := contractAbi.Events["TaskCreated"]
		for _, vLog := range logs {
			data := vLog.Data

			eventArgs, err := event.Inputs.Unpack(data)
			if err != nil {
				o.logger.Error("Failed to unpack event inputs", "err", err)
				return err
			}
			if eventArgs != nil {
				taskResponse := o.ProcessNewTaskCreatedLog(eventArgs)
				sig, resBytes, err := o.SignTaskResponse(taskResponse)
				if err != nil {
					o.logger.Error("Failed to sign task response", "err", err)
					continue
				}
				taskInfo, _ := o.avsReader.GetTaskInfo(&bind.CallOpts{}, o.avsAddr.String(), taskResponse.TaskID)
				go o.SendSignedTaskResponseToExocore(context.Background(), taskResponse.TaskID, resBytes, sig, taskInfo)
			}
		}
	}
	return nil
}

// ProcessNewTaskCreatedLog TaskResponse is the struct that is signed and sent to the exocore as a task response.
func (o *Operator) ProcessNewTaskCreatedLog(eventArgs []interface{}) *core.TaskResponse {
	o.logger.Debug("Received new task", "task", eventArgs)
	o.logger.Info("Received new task",
		"id", eventArgs[0].(*big.Int),
		"name", eventArgs[2].(string),
	)
	taskId := eventArgs[0].(*big.Int)
	taskResponse := &core.TaskResponse{
		TaskID:    taskId.Uint64(),
		NumberSum: taskId,
	}
	return taskResponse
}

func (o *Operator) SignTaskResponse(taskResponse *core.TaskResponse) ([]byte, []byte, error) {

	taskResponseHash, data, err := core.GetTaskResponseDigestEncodeByjson(*taskResponse)
	if err != nil {
		o.logger.Error("Error SignTaskResponse with getting task response header hash. skipping task (this is not expected and should be investigated)", "err", err)
		return nil, nil, err
	}
	msgBytes := taskResponseHash[:]

	sig := o.blsKeypair.Sign(msgBytes)

	return sig.Marshal(), data, nil
}

func (o *Operator) SendSignedTaskResponseToExocore(
	ctx context.Context,
	taskId uint64,
	taskResponse []byte,
	blsSignature []byte,
	taskInfo []uint64) (string, error) {

	startingEpoch := taskInfo[0]
	taskResponsePeriod := taskInfo[1]
	taskStatisticalPeriod := taskInfo[2]

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err() // Gracefully exit if context is canceled
		default:
			// Fetch the current epoch information
			epochIdentifier, err := o.avsReader.GetAVSEpochIdentifier(&bind.CallOpts{}, o.avsAddr.String())
			if err != nil {
				o.logger.Error("Cannot GetAVSEpochIdentifier", "err", err)
				return "", fmt.Errorf("failed to get AVS info: %w", err) // Stop on persistent error
			}

			num, err := o.avsReader.GetCurrentEpoch(&bind.CallOpts{}, epochIdentifier)
			if err != nil {
				o.logger.Error("Cannot exec GetCurrentEpoch", "err", err)
				return "", fmt.Errorf("failed to get current epoch: %w", err) // Stop on persistent error
			}

			currentEpoch := uint64(num)
			o.logger.Info("current epoch  is :", "currentEpoch", currentEpoch)
			if currentEpoch > startingEpoch+taskResponsePeriod+taskStatisticalPeriod {
				o.logger.Info("Exiting loop: Task period has passed")
				return "Task period has passed", nil
			}

			switch {
			case currentEpoch <= startingEpoch:
				o.logger.Info("current epoch  is less than starting epoch", "currentEpoch", currentEpoch, "startingEpoch", startingEpoch)

			case currentEpoch <= startingEpoch+taskResponsePeriod:
				o.logger.Info("Execute Phase One Submission Task", "currentEpoch", currentEpoch,
					"startingEpoch", startingEpoch, "taskResponsePeriod", taskResponsePeriod)
				o.logger.Info("Submitting task response for task response period",
					"taskAddr", o.avsAddr.String(), "taskId", taskId, "operator-addr", o.operatorAddr)
				_, err := o.avsWriter.OperatorSubmitTask(
					ctx,
					taskId,
					nil,
					blsSignature,
					o.avsAddr.String(),
					1)
				if err != nil {
					o.logger.Error("Avs failed to OperatorSubmitTask", "err", err)
					return "", fmt.Errorf("failed to submit task during taskResponsePeriod: %w", err)
				}

			case currentEpoch <= startingEpoch+taskResponsePeriod+taskStatisticalPeriod && currentEpoch > startingEpoch+taskResponsePeriod:
				o.logger.Info("Execute Phase Two Submission Task", "currentEpoch", currentEpoch,
					"startingEpoch", startingEpoch, "taskResponsePeriod", taskResponsePeriod, "taskStatisticalPeriod", taskStatisticalPeriod)
				o.logger.Info("Submitting task response for statistical period",
					"taskAddr", o.avsAddr.String(), "taskId", taskId, "operator-addr", o.operatorAddr)
				_, err := o.avsWriter.OperatorSubmitTask(
					ctx,
					taskId,
					taskResponse,
					blsSignature,
					o.avsAddr.String(),
					2)
				if err != nil {
					o.logger.Error("Avs failed to OperatorSubmitTask", "err", err)
					return "", fmt.Errorf("failed to submit task during statistical period: %w", err)
				}

			default:
				o.logger.Info("Current epoch is not within expected range", "currentEpoch", currentEpoch)
				return "", fmt.Errorf("current epoch %d is not within expected range %d", currentEpoch, startingEpoch)
			}
			// Disable the uselesss ticker.
			// // Dynamic sleep based on the epoch identifier
			// var sleepDuration time.Duration
			// switch epochIdentifier {
			// case DayEpochID:
			// 	sleepDuration = 24 * time.Hour
			// case HourEpochID:
			// 	sleepDuration = time.Hour
			// case MinuteEpochID:
			// 	sleepDuration = time.Minute
			// case WeekEpochID:
			// 	sleepDuration = 7 * 24 * time.Hour
			// default:
			// 	o.logger.Warn("Unknown epoch identifier", "epochIdentifier", epochIdentifier)
			// 	sleepDuration = time.Minute // Default to a safe short duration
			// }

			// // Use a ticker instead of time.Sleep for more control and graceful handling
			// ticker := time.NewTicker(sleepDuration)
			// defer ticker.Stop() // Ensure ticker stops even if the function exits early
			// select {
			// case <-ticker.C:
			// 	continue // Proceed to the next iteration
			// case <-ctx.Done():
			// 	return "", ctx.Err() // Handle cancellation
			// }
		}
	}
}

func SwitchEthAddressToExoAddress(ethAddress string) (string, error) {
	b, err := hex.DecodeString(ethAddress[2:])
	if err != nil {
		return "", fmt.Errorf("failed to decode eth address: %w", err)
	}

	// Generate exo address
	HRP := "exo"
	exoAddress, err := bech32.EncodeFromBase256(HRP, b)
	if err != nil {
		return "", fmt.Errorf("failed to encode exo address: %w", err)
	}

	return exoAddress, nil
}
