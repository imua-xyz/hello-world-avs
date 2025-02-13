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
	blscommon "github.com/prysmaticlabs/prysm/v5/crypto/bls/common"
	"math/big"
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
	config         types.NodeConfig
	logger         sdklogging.Logger
	ethClient      eth.EthClient
	nodeApi        *nodeapi.NodeApi
	avsWriters     []chain.ExoWriter
	avsReaders     []chain.ExoChainReader
	avsSubscribers []chain.AvsRegistrySubscriber

	blsKeypairs  []blscommon.SecretKey
	operatorAddr common.Address
	// receive new tasks in this chan (typically from listening to onchain event)
	newTaskCreatedChan chan *avs.ContracthelloWorldTaskCreated
	// needed when opting in to avs (allow this service manager contract to slash operator)
	avsAddresses     []common.Address
	epochIdentifiers []string
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

	var blsKeyPairs []blscommon.SecretKey
	for _, path := range c.BlsPrivateKeyStorePath {
		blsKeyPair, err := bls.ReadPrivateKeyFromFile(path, "")
		if err != nil {
			logger.Error("Cannot parse bls private key", "err", err)
			return nil, err
		}
		blsKeyPairs = append(blsKeyPairs, blsKeyPair)
	}

	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	ecdsaKeyPassword := ""
	signerV2, operatorSender, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: c.OperatorEcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil {
		panic(err)
	}
	logger.Info("operatorSender:", "operatorSender", operatorSender.String())

	balance, err := ethRpcClient.BalanceAt(context.Background(), operatorSender, nil)
	if err != nil {
		logger.Error("Cannot get Balance", "err", err)
	}
	if balance.Cmp(big.NewInt(0)) != 1 {
		logger.Error("operatorSender has not enough Balance")
	}
	if c.OperatorAddress != operatorSender.String() {
		logger.Error("operatorSender is not equal OperatorAddress")
	}
	txMgr := txmgr.NewSimpleTxManager(ethRpcClient, logger, signerV2, common.HexToAddress(c.OperatorAddress))
	var hexes []common.Address
	var avsReaders []chain.ExoChainReader
	var avsWriters []chain.ExoWriter
	var avsSubscribers []chain.AvsRegistrySubscriber
	var epochIdentifiers []string
	for _, strAddress := range c.AVSAddresses {
		hexAddress := common.HexToAddress(strAddress)
		hexes = append(hexes, hexAddress)
		avsReader, _ := chain.BuildExoChainReader(
			common.HexToAddress(strAddress),
			ethRpcClient,
			logger)
		avsWriter, _ := chain.BuildExoChainWriter(
			common.HexToAddress(strAddress),
			ethRpcClient,
			logger,
			txMgr)

		avsSubscriber, _ := chain.BuildAvsRegistryChainSubscriber(
			common.HexToAddress(strAddress),
			ethWsClient,
			logger,
		)
		avsReaders = append(avsReaders, *avsReader)
		avsWriters = append(avsWriters, avsWriter)
		avsSubscribers = append(avsSubscribers, avsSubscriber)

		epochIdentifier, err := avsReader.GetAVSEpochIdentifier(&bind.CallOpts{}, strAddress)
		if err != nil {
			logger.Error("Cannot GetAVSEpochIdentifier", "err", err)
			return nil, err
		}
		epochIdentifiers = append(epochIdentifiers, epochIdentifier)
	}

	operator := &Operator{
		config:             c,
		logger:             logger,
		nodeApi:            nodeApi,
		ethClient:          ethRpcClient,
		avsWriters:         avsWriters,
		avsReaders:         avsReaders,
		avsSubscribers:     avsSubscribers,
		blsKeypairs:        blsKeyPairs,
		operatorAddr:       common.HexToAddress(c.OperatorAddress),
		newTaskCreatedChan: make(chan *avs.ContracthelloWorldTaskCreated),
		avsAddresses:       hexes,
		epochIdentifiers:   epochIdentifiers,
	}

	if c.RegisterOperatorOnStartup {

		if err != nil {
			return nil, err
		}
		operator.registerOperatorOnStartup()
	}

	logger.Info("Operator info", "operatorAddr", c.OperatorAddress)

	return operator, nil
}

func (o *Operator) Start(ctx context.Context) error {
	// 1.operator register exocore
	// 2.operator opt-in avs
	// 3.operator accept staker delegation so that avs voting power is not 0, otherwise the task cannot be created
	// 4.operator register BLSPublicKey
	// 5.operator submit task
	for i, avsWriter := range o.avsWriters {
		// operator register BLSPublicKey  via evm tx
		msgBytes := crypto.Keccak256Hash([]byte("hello-avs")).Bytes()
		sig := o.blsKeypairs[i].Sign(msgBytes)
		_, err := avsWriter.RegisterBLSPublicKey(
			context.Background(),
			o.avsAddresses[i].String(),
			o.blsKeypairs[i].PublicKey().Marshal(),
			sig.Marshal(),
			msgBytes)
		if err != nil {
			o.logger.Error("operator failed to registerBLSPublicKey", "err", err)
			return err
		}
	}
	//depoist and delegate
	err := o.Deposit()
	if err != nil {
		o.logger.Error("Cannot Deposit", "err", err)
		return err
	}
	err = o.Delegate()
	if err != nil {
		o.logger.Error("Cannot Delegate", "err", err)
		return err
	}
	err = o.SelfDelegate()
	if err != nil {
		o.logger.Error("Cannot SelfDelegate", "err", err)
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
	o.logger.Info("Event firstHeight: ", "First detected block height", firstHeight)
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
			// o.logger.Info("newHeight", newHeight)
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
				taskInfo, _ := o.avsReaders.GetTaskInfo(&bind.CallOpts{}, o.avsAddr.String(), taskResponse.TaskID)
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
		"numberToBeSquared", eventArgs[3].(uint64),
	)
	taskId := eventArgs[0].(*big.Int)
	numberToBeSquared := eventArgs[3].(uint64)
	taskResponse := &core.TaskResponse{
		TaskID:            taskId.Uint64(),
		NumberToBeSquared: numberToBeSquared * numberToBeSquared,
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

	sig := o.blsKeypairs.Sign(msgBytes)

	return sig.Marshal(), data, nil
}

func (o *Operator) SendSignedTaskResponseToExocore(
	ctx context.Context,
	taskId uint64,
	taskResponse []byte,
	blsSignature []byte,
	taskInfo avs.TaskInfo) (string, error) {

	startingEpoch := taskInfo.StartingEpoch
	taskResponsePeriod := taskInfo.TaskResponsePeriod
	taskStatisticalPeriod := taskInfo.TaskStatisticalPeriod

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err() // Gracefully exit if context is canceled
		default:
			// Fetch the current epoch information
			epochIdentifier, err := o.avsReaders.GetAVSEpochIdentifier(&bind.CallOpts{}, o.avsAddr.String())
			if err != nil {
				o.logger.Error("Cannot GetAVSEpochIdentifier", "err", err)
				return "", fmt.Errorf("failed to get AVS info: %w", err) // Stop on persistent error
			}

			num, err := o.avsReaders.GetCurrentEpoch(&bind.CallOpts{}, epochIdentifier)
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
				_, err := o.avsWriters.OperatorSubmitTask(
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
				_, err := o.avsWriters.OperatorSubmitTask(
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
