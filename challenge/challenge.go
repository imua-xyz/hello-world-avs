package challenge

import (
	"context"
	"fmt"
	avs "github.com/ExocoreNetwork/exocore-avs/contracts/bindings/avs"
	chain "github.com/ExocoreNetwork/exocore-avs/core/chainio"
	"github.com/ExocoreNetwork/exocore-avs/core/chainio/eth"
	"github.com/ExocoreNetwork/exocore-avs/types"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	sdklogging "github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ExocoreNetwork/exocore-sdk/nodeapi"
	"github.com/ExocoreNetwork/exocore-sdk/signerv2"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"os"
	"time"
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

type Challenger struct {
	config        types.NodeConfig
	logger        sdklogging.Logger
	ethClient     eth.EthClient
	nodeApi       *nodeapi.NodeApi
	avsWriter     chain.ExoWriter
	avsReader     chain.ExoChainReader
	avsSubscriber chain.AvsRegistrySubscriber
	// receive new tasks in this chan (typically from listening to onchain event)
	newTaskCreatedChan chan *avs.ContracthelloWorldTaskCreated
	// needed when opting in to avs (allow this service manager contract to slash operator)
	avsAddr         common.Address
	epochIdentifier string
}

func NewChallengeFromConfig(c types.NodeConfig) (*Challenger, error) {
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

	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("AVS_ECDSA_KEY_PASSWORD")
	if !ok {
		logger.Info("AVS_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	signerV2, challengeSender, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: c.AVSEcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil {
		panic(err)
	}
	logger.Info("challengeSender:", "challengeSender", challengeSender.String())

	balance, err := ethRpcClient.BalanceAt(context.Background(), challengeSender, nil)
	if err != nil {
		logger.Error("Cannot get Balance", "err", err)
	}
	if balance.Cmp(big.NewInt(0)) != 1 {
		logger.Error("challengeSender has not enough Balance")
	}
	if c.AVSOwnerAddress != challengeSender.String() {
		logger.Error("challengeSender is not equal AVSOwnerAddress")
	}
	txMgr := txmgr.NewSimpleTxManager(ethRpcClient, logger, signerV2, common.HexToAddress(c.AVSOwnerAddress))

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
	challenger := &Challenger{
		config:             c,
		logger:             logger,
		nodeApi:            nodeApi,
		ethClient:          ethRpcClient,
		avsWriter:          avsWriter,
		avsReader:          *avsReader,
		avsSubscriber:      avsSubscriber,
		newTaskCreatedChan: make(chan *avs.ContracthelloWorldTaskCreated),
		avsAddr:            common.HexToAddress(c.AVSAddress),
		epochIdentifier:    epochIdentifier,
	}
	logger.Info("challenger info", "challengeAddr", c.AVSOwnerAddress)

	return challenger, nil
}
func (o *Challenger) Exec(ctx context.Context) error {
	taskInfo, _ := o.avsReader.GetTaskInfo(&bind.CallOpts{}, o.avsAddr.String(), 1)
	infos, _ := o.avsReader.GetOperatorTaskResponseList(&bind.CallOpts{}, taskInfo.TaskContractAddress.String(), taskInfo.TaskID)
	task := &avs.AvsServiceContractChallengeReq{
		TaskId:            taskInfo.TaskID,
		TaskAddress:       taskInfo.TaskContractAddress,
		NumberToBeSquared: 350,
		Infos:             infos,
		SignedOperators:   taskInfo.SignedOperators,
		NoSignedOperators: taskInfo.NoSignedOperators,
		TaskTotalPower:    taskInfo.TaskTotalPower,
	}
	o.logger.Info("challenger info", "challenge-TaskResponse", infos[0].TaskResponse)
	_, err := o.avsWriter.Challenge(
		ctx,
		*task)

	if err != nil {
		o.logger.Error("Challeger failed to raiseAndResolveChallenge", "err", err)
		return fmt.Errorf("failed to raiseAndResolveChallenge: %w", err)
	}

	return nil
}
func (o *Challenger) Start(ctx context.Context) error {
	// 1. First, the task reaches the challenge period and the module is verified
	// 2. Is an effective task:
	// Has not been processed, query kv to make sure it has not been processed module validation
	// At the same time, optInOperators is not empty, which is ensured when creating, and avs usd will be required not to be 0 module verification when creating
	// 3. Get the array string optInOperators, iterate through it, query and verify that the submitted task respond of each Operator is equal to the formula,
	// If you validate the total passed power by putting it into the array string[] eligibleRewardOperators and getting the data from operatorActivePower, Do not verify by putting in the array string[] eligibleSlashOperators;
	// 4. Calculate the proportion of passed power in the total taskTotalPower
	// 5. Compare whether the above proportion is greater than or equal to thresholdPercentage. If it is greater than, isExpected is true; otherwise, it is false
	o.logger.Infof("Starting Challenge.")
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
func (o *Challenger) GetLog(height int64) error {
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
				task := o.ProcessNewTaskCreatedLog(eventArgs)

				taskInfo, err := o.avsReader.GetTaskInfo(&bind.CallOpts{}, o.avsAddr.String(), task.TaskId)
				if err != nil {
					o.logger.Error("Failed to GetTaskInfo", "err", err)
					return err
				}
				go func() {
					_, err := o.TriggerChallege(context.Background(), *task, taskInfo)
					if err != nil {

					}
				}()
			}
		}
	}
	return nil
}

// ProcessNewTaskCreatedLog TaskResponse is the struct that is signed and sent to the exocore as a task response.
func (o *Challenger) ProcessNewTaskCreatedLog(eventArgs []interface{}) *avs.AvsServiceContractChallengeReq {
	o.logger.Debug("Received new task", "task", eventArgs)
	o.logger.Info("Received new task",
		"id", eventArgs[0].(*big.Int),
		"name", eventArgs[2].(string),
		"numberToBeSquared", eventArgs[3].(uint64),
	)

	task := &avs.AvsServiceContractChallengeReq{
		TaskId:            eventArgs[0].(*big.Int).Uint64(),
		NumberToBeSquared: eventArgs[3].(uint64),
	}
	return task
}

func (o *Challenger) TriggerChallege(
	ctx context.Context,
	task avs.AvsServiceContractChallengeReq,
	taskInfo avs.TaskInfo) (string, error) {
	o.logger.Debug("TriggerChallege", "taskInfo", taskInfo)
	epochIdentifier, err := o.avsReader.GetAVSEpochIdentifier(&bind.CallOpts{}, o.avsAddr.String())
	startingEpoch := taskInfo.StartingEpoch
	taskResponsePeriod := taskInfo.TaskResponsePeriod
	taskStatisticalPeriod := taskInfo.TaskStatisticalPeriod
	num := startingEpoch + taskResponsePeriod + taskStatisticalPeriod

	var sleepDuration time.Duration
	switch epochIdentifier {
	case DayEpochID:
		sleepDuration = time.Duration(num) * 24 * time.Hour
	case HourEpochID:
		sleepDuration = time.Duration(num) * time.Hour
	case MinuteEpochID:
		sleepDuration = time.Duration(num) * time.Minute
	case WeekEpochID:
		sleepDuration = time.Duration(num) * 7 * 24 * time.Hour
	default:
		o.logger.Info("Unknown epoch identifier", "epochIdentifier", epochIdentifier)
		sleepDuration = time.Minute // Default to a safe short duration
	}

	time.Sleep(sleepDuration)

	if taskInfo.IsExpected {
		o.logger.Infof("Task %d is expected. Skipping challenge", task.TaskId)
		return "", nil
	}
	if len(taskInfo.OptInOperators) < 1 {
		o.logger.Infof("Task %d does not have any optIn operators. Skipping challenge", task.TaskId)
		return "", nil
	}

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err() // Gracefully exit if context is canceled
		default:
			// Fetch the current epoch information
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
			// o.logger.Info("current epoch  is :", "currentEpoch", currentEpoch)
			if currentEpoch > startingEpoch+taskResponsePeriod+taskStatisticalPeriod {
				o.logger.Info("Execute raiseAndResolveChallenge", "currentEpoch", currentEpoch,
					"startingEpoch", startingEpoch, "taskResponsePeriod", taskResponsePeriod, "taskStatisticalPeriod", taskStatisticalPeriod)
				_, err := o.avsWriter.Challenge(
					ctx,
					task)
				if err != nil {
					o.logger.Error("Challeger failed to raiseAndResolveChallenge", "err", err)
					return "", fmt.Errorf("failed to raiseAndResolveChallenge: %w", err)
				}

			}
		}
	}
}
