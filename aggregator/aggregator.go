package aggregator

import (
	"context"
	"github.com/ExocoreNetwork/exocore-avs/aggregator/types"
	chain "github.com/ExocoreNetwork/exocore-avs/core/chainio"
	"github.com/ExocoreNetwork/exocore-avs/core/chainio/eth"
	"github.com/ExocoreNetwork/exocore-avs/core/config"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	sdklogging "github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ExocoreNetwork/exocore-sdk/signerv2"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"os"
	"time"
)

const (
	taskChallengeWindowBlock = 100
	blockTime                = 12 * time.Second
	avsName                  = "avs-demo-numToSquare"
)

// Aggregator currently we only use a single quorum of  token,A task can only be responded onchain by having enough operators sign on it such that their stake in each quorum reaches the QuorumThresholdPercentage.
type Aggregator struct {
	logger    logging.Logger
	avsWriter chain.EXOWriter
}

// NewAggregator creates a new Aggregator with the provided config.
func NewAggregator(c *config.Config) (*Aggregator, error) {
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

	ethRpcClient, err := eth.NewClient(c.EthHttpRpcUrl)
	if err != nil {
		logger.Errorf("Cannot create http ethclient", "err", err)
		return nil, err
	}
	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: c.EcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil {
		panic(err)
	}

	txMgr := txmgr.NewSimpleTxManager(ethRpcClient, logger, signerV2, common.HexToAddress(c.AggregatorAddress.String()))
	avsWriter, _ := chain.BuildELChainWriter(
		common.HexToAddress(c.AggregatorAddress.String()),
		ethRpcClient,
		logger,
		txMgr)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	return &Aggregator{
		logger:    c.Logger,
		avsWriter: avsWriter,
	}, nil
}

func (agg *Aggregator) Start(ctx context.Context) error {
	agg.logger.Infof("Starting aggregator.")
	agg.logger.Infof("Starting aggregator rpc server.")
	ticker := time.NewTicker(20 * time.Second)
	agg.logger.Infof("Aggregator set to send new task every 10 seconds...")
	defer ticker.Stop()
	taskNum := int64(0)
	// send the first task
	_ = agg.sendNewTask(big.NewInt(taskNum), big.NewInt(taskNum))
	taskNum++

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			agg.logger.Info("sendNewTask-num", taskNum)
			err := agg.sendNewTask(big.NewInt(taskNum), big.NewInt(taskNum))
			taskNum++
			if err != nil {
				// we log the errors inside sendNewTask() so here we just continue to the next task
				continue
			}
		}
	}
}

// sendNewTask sends a new task to the task manager contract, and updates the Task dict struct
// with the information of operators opted into quorum 0 at the block of task creation.
func (agg *Aggregator) sendNewTask(numOneToSum, numTwoToSum *big.Int) error {
	agg.logger.Info("Aggregator sending new task", "numberOneToSum:", numOneToSum, "numberTwoToSum:", numTwoToSum)
	_, err := agg.avsWriter.CreateNewTask(
		context.Background(),
		"",
		types.TaskResponsePeriod,
		types.TaskChallengePeriod,
		types.ThresholdPercentage,
		types.TaskStatisticalPeriod)

	if err != nil {
		agg.logger.Error("Aggregator failed to send number to square", "err", err)
		return err
	}

	return nil
}
