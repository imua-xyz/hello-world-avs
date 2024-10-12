package aggregator

import (
	"context"
	"github.com/ExocoreNetwork/exocore-avs/aggregator/types"
	cstaskmanager "github.com/ExocoreNetwork/exocore-avs/bindings/AvsTaskManager"
	"github.com/ExocoreNetwork/exocore-avs/core"
	"github.com/ExocoreNetwork/exocore-avs/core/chainio"
	"github.com/ExocoreNetwork/exocore-avs/core/config"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients"
	sdkclients "github.com/ExocoreNetwork/exocore-sdk/chainio/clients"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ExocoreNetwork/exocore-sdk/services/avsregistry"
	blsagg "github.com/ExocoreNetwork/exocore-sdk/services/bls_aggregation"
	oppubkeysserv "github.com/ExocoreNetwork/exocore-sdk/services/operatorpubkeys"
	sdktypes "github.com/ExocoreNetwork/exocore-sdk/types"
	"math/big"
	"sync"
	"time"
)

const (
	taskChallengeWindowBlock = 100
	blockTime                = 12 * time.Second
	avsName                  = "avs-demo-numToSquare"
)

// Aggregator currently we only use a single quorum of  token,A task can only be responded onchain by having enough operators sign on it such that their stake in each quorum reaches the QuorumThresholdPercentage.
type Aggregator struct {
	logger           logging.Logger
	serverIpPortAddr string
	avsWriter        chainio.AvsWriterer
	// aggregation related fields
	blsAggregationService blsagg.BlsAggregationService
	tasks                 map[types.TaskIndex]cstaskmanager.IIncredibleSquaringTaskManagerTask
	tasksMu               sync.RWMutex
	taskResponses         map[types.TaskIndex]map[sdktypes.TaskResponseDigest]cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
	taskResponsesMu       sync.RWMutex
}

// NewAggregator creates a new Aggregator with the provided config.
func NewAggregator(c *config.Config) (*Aggregator, error) {

	avsReader, err := chainio.BuildAvsReaderFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create avsReader", "err", err)
		return nil, err
	}

	avsWriter, err := chainio.BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	chainConfig := sdkclients.BuildAllConfig{
		EthHttpUrl:                 c.EthHttpRpcUrl,
		EthWsUrl:                   c.EthWsRpcUrl,
		RegistryCoordinatorAddr:    c.AvsRegistryCoordinatorAddr.String(),
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddr.String(),
		AvsName:                    avsName,
	}
	clients, err := clients.BuildAll(chainConfig, c.AggregatorAddress, c.SignerFn, c.Logger)
	if err != nil {
		c.Logger.Errorf("Cannot create sdk clients", "err", err)
		return nil, err
	}

	operatorPubkeysService := oppubkeysserv.NewOperatorPubkeysServiceInMemory(context.Background(), clients.AvsRegistryChainSubscriber, clients.AvsRegistryChainReader, c.Logger)
	avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(avsReader, operatorPubkeysService, c.Logger)
	blsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, c.Logger)

	return &Aggregator{
		logger:                c.Logger,
		serverIpPortAddr:      c.AggregatorServerIpPortAddr,
		avsWriter:             avsWriter,
		blsAggregationService: blsAggregationService,
		tasks:                 make(map[types.TaskIndex]cstaskmanager.IIncredibleSquaringTaskManagerTask),
		taskResponses:         make(map[types.TaskIndex]map[sdktypes.TaskResponseDigest]cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse),
	}, nil
}

func (agg *Aggregator) Start(ctx context.Context) error {
	agg.logger.Infof("Starting aggregator.")
	agg.logger.Infof("Starting aggregator rpc server.")
	go agg.startServer(ctx)
	ticker := time.NewTicker(20 * time.Second)
	agg.logger.Infof("Aggregator set to send new task every 10 seconds...")
	defer ticker.Stop()
	taskNum := int64(0)
	// send the first task
	_ = agg.sendNewTask(big.NewInt(taskNum))
	taskNum++

	for {
		select {
		case <-ctx.Done():
			return nil
		case blsAggServiceResp := <-agg.blsAggregationService.GetResponseChannel():
			agg.logger.Info("Received response from blsAggregationService", "blsAggServiceResp", blsAggServiceResp)
			agg.sendAggregatedResponseToContract(blsAggServiceResp)
		case <-ticker.C:
			err := agg.sendNewTask(big.NewInt(taskNum))
			taskNum++
			if err != nil {
				// we log the errors inside sendNewTask() so here we just continue to the next task
				continue
			}
		}
	}
}

func (agg *Aggregator) sendAggregatedResponseToContract(blsAggServiceResp blsagg.BlsAggregationServiceResponse) {
	// TODO: check if blsAggServiceResp contains an err
	if blsAggServiceResp.Err != nil {
		agg.logger.Error("BlsAggregationServiceResponse contains an error", "err", blsAggServiceResp.Err)
		// panicing to help with debugging (fail fast), but we shouldn't panic if we run this in production
		panic(blsAggServiceResp.Err)
	}
	var nonSignerPubkeys []cstaskmanager.BN254G1Point
	for _, nonSignerPubkey := range blsAggServiceResp.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, core.ConvertToBN254G1Point(nonSignerPubkey))
	}
	var quorumApks []cstaskmanager.BN254G1Point
	for _, quorumApk := range blsAggServiceResp.QuorumApksG1 {
		quorumApks = append(quorumApks, core.ConvertToBN254G1Point(quorumApk))
	}
	nonSignerStakesAndSignature := cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature{
		NonSignerPubkeys:             nonSignerPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        core.ConvertToBN254G2Point(blsAggServiceResp.SignersApkG2),
		Sigma:                        core.ConvertToBN254G1Point(blsAggServiceResp.SignersAggSigG1.G1Point),
		NonSignerQuorumBitmapIndices: blsAggServiceResp.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             blsAggServiceResp.QuorumApkIndices,
		TotalStakeIndices:            blsAggServiceResp.TotalStakeIndices,
		NonSignerStakeIndices:        blsAggServiceResp.NonSignerStakeIndices,
	}

	agg.logger.Info("Threshold reached. Sending aggregated response onchain.",
		"taskIndex", blsAggServiceResp.TaskIndex,
	)
	agg.tasksMu.RLock()
	task := agg.tasks[blsAggServiceResp.TaskIndex]
	agg.tasksMu.RUnlock()
	agg.taskResponsesMu.RLock()
	taskResponse := agg.taskResponses[blsAggServiceResp.TaskIndex][blsAggServiceResp.TaskResponseDigest]
	agg.taskResponsesMu.RUnlock()
	_, err := agg.avsWriter.SendAggregatedResponse(context.Background(), task, taskResponse, nonSignerStakesAndSignature)
	if err != nil {
		agg.logger.Error("Aggregator failed to respond to task", "err", err)
	}
}

// sendNewTask sends a new task to the task manager contract, and updates the Task dict struct
// with the information of operators opted into quorum 0 at the block of task creation.
func (agg *Aggregator) sendNewTask(numOneToSum, numTwoToSum *big.Int) error {
	agg.logger.Info("Aggregator sending new task", "numberOneToSum:", numOneToSum, "numberTwoToSum:", numTwoToSum)
	// Send number to sum to the task manager contract
	newTask, taskIndex, err := agg.avsWriter.SendNewTaskNumberToSum(context.Background(), numOneToSum, numTwoToSum, types.ThresholdNumerator)
	if err != nil {
		agg.logger.Error("Aggregator failed to send number to square", "err", err)
		return err
	}

	agg.tasksMu.Lock()
	agg.tasks[taskIndex] = newTask
	agg.tasksMu.Unlock()

	quorumThresholdPercentages := make([]uint32, len(newTask.QuorumNumbers))
	for i := range newTask.QuorumNumbers {
		quorumThresholdPercentages[i] = newTask.QuorumThresholdPercentage
	}
	taskTimeToExpiry := taskChallengeWindowBlock * blockTime
	agg.blsAggregationService.InitializeNewTask(taskIndex, newTask.TaskCreatedBlock, newTask.QuorumNumbers, quorumThresholdPercentages, taskTimeToExpiry)
	return nil
}
