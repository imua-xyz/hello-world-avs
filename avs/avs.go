package avs

import (
	"context"
	sdkmath "cosmossdk.io/math"
	avs "github.com/ExocoreNetwork/exocore-avs/contracts/bindings/avs"
	"github.com/ExocoreNetwork/exocore-avs/core"
	chain "github.com/ExocoreNetwork/exocore-avs/core/chainio"
	"github.com/ExocoreNetwork/exocore-avs/core/chainio/eth"
	"github.com/ExocoreNetwork/exocore-avs/types"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	sdkEcdsa "github.com/ExocoreNetwork/exocore-sdk/crypto/ecdsa"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	sdklogging "github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ExocoreNetwork/exocore-sdk/signerv2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/rand"
	"os"
	"time"
)

const (
	avsName    = "hello-avs-demo"
	maxRetries = 10
	retryDelay = 5 * time.Second
)

type Avs struct {
	logger                logging.Logger
	avsWriter             chain.ExoWriter
	avsReader             chain.ExoReader
	avsAddress            string
	createTaskInterval    int64
	taskResponsePeriod    uint64
	taskChallengePeriod   uint64
	thresholdPercentage   uint64
	taskStatisticalPeriod uint64
}

// NewAvs creates a new Avs with the provided config.
func NewAvs(c *types.NodeConfig) (*Avs, error) {
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

	ethRpcClient, err := eth.NewClient(c.EthRpcUrl)
	if err != nil {
		logger.Error("Cannot create http ethclient", "err", err)
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

	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: c.AVSEcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil {
		panic(err)
	}

	if c.AVSAddress == "" {
		logger.Info("AVS_ADDRESS env var not set. will deploy avs contract")

		key, err := sdkEcdsa.ReadKey(c.AVSEcdsaPrivateKeyStorePath, ecdsaKeyPassword)

		avsAddr, _, err := chain.DeployAVS(
			ethRpcClient,
			logger,
			*key,
			chainId,
		)
		if err != nil {
			panic(err)
		}

		c.AVSAddress = avsAddr.String()
		c.TaskAddress = avsAddr.String()
		c.AVSRewardAddress = avsAddr.String()
		c.AVSSlashAddress = avsAddr.String()
		filePath, err := core.GetFileInCurrentDirectory("config.yaml")
		if err != nil {
			panic(err)
		}
		err = core.UpdateYAMLWithComments(filePath, "avs_address", avsAddr.String())
		err = core.UpdateYAMLWithComments(filePath, "avs_reward_address", avsAddr.String())
		err = core.UpdateYAMLWithComments(filePath, "avs_slash_address", avsAddr.String())
		err = core.UpdateYAMLWithComments(filePath, "task_address", avsAddr.String())

		if err != nil {
			logger.Error("AVS_ADDRESS env var not set. will deploy avs contract")

		}
	}

	txMgr := txmgr.NewSimpleTxManager(ethRpcClient, logger, signerV2, common.HexToAddress(c.AVSOwnerAddress))
	avsWriter, err := chain.BuildExoChainWriter(
		common.HexToAddress(c.AVSAddress),
		ethRpcClient,
		logger,
		txMgr)
	if err != nil {
		logger.Error("Cannot create avsWriter", "err", err)
		return nil, err
	}

	avsReader, err := chain.BuildExoChainReader(
		common.HexToAddress(c.AVSAddress),
		ethRpcClient,
		logger)
	if err != nil {
		logger.Error("Cannot create exoChainReader", "err", err)
		return nil, err
	}
	time.Sleep(retryDelay)
	info, err := avsReader.GetAVSEpochIdentifier(&bind.CallOpts{}, c.AVSAddress)
	if err != nil {
		logger.Error("Cannot GetAVSEpochIdentifier", "err", err)
		return nil, err
	}
	if info == "" {
		params := avs.AVSParams{
			Sender:              common.HexToAddress(c.AVSOwnerAddress),
			AvsName:             avsName,
			MinStakeAmount:      c.MinStakeAmount,
			TaskAddr:            common.HexToAddress(c.TaskAddress),
			SlashAddr:           common.HexToAddress(c.AVSRewardAddress),
			RewardAddr:          common.HexToAddress(c.AVSSlashAddress),
			AvsOwnerAddress:     core.ConvertToEthAddresses(c.AvsOwnerAddresses),
			WhitelistAddress:    core.ConvertToEthAddresses(c.WhitelistAddresses),
			AssetIds:            c.AssetIds,
			AvsUnbondingPeriod:  c.AvsUnbondingPeriod,
			MinSelfDelegation:   c.MinSelfDelegation,
			EpochIdentifier:     c.EpochIdentifier,
			MiniOptInOperators:  1,
			MinTotalStakeAmount: 1,
			AvsRewardProportion: 5,
			AvsSlashProportion:  5,
		}
		_, err = avsWriter.RegisterAVSToExocore(context.Background(),
			params,
		)
		if err != nil {
			logger.Error("register Avs failed ", "err", err)
			return &Avs{}, err
		}
	}

	return &Avs{
		logger:                logger,
		avsWriter:             avsWriter,
		avsReader:             avsReader,
		avsAddress:            c.AVSAddress,
		createTaskInterval:    c.CreateTaskInterval,
		taskResponsePeriod:    c.TaskResponsePeriod,
		taskChallengePeriod:   c.TaskChallengePeriod,
		thresholdPercentage:   c.ThresholdPercentage,
		taskStatisticalPeriod: c.TaskStatisticalPeriod,
	}, nil
}

func (avs *Avs) Start(ctx context.Context) error {
	avs.logger.Infof("Starting avs.")
	ticker := time.NewTicker(time.Duration(avs.createTaskInterval) * time.Second)
	avs.logger.Infof("Avs owner set to send new task every %d seconds", avs.createTaskInterval)
	defer ticker.Stop()
	taskNum := int64(1)
	// send the first task
	err := avs.sendNewTask()
	if err != nil {
		// we log the errors inside sendNewTask() so here we just continue to do the next task
		avs.logger.Info("sendNewTask encountered an error: %v; continuing to do the next task.", err)
	}
	taskNum++
	for {
		select {
		case <-ctx.Done():
			avs.logger.Info("Context canceled; stopping AVS.")
			return nil
		case <-ticker.C:
			avs.logger.Info("sendNewTask-num:", "taskNum", taskNum)
			err := avs.sendNewTask()
			if err != nil {
				// we log the errors inside sendNewTask() so here we just continue to do the next task
				avs.logger.Info("sendNewTask encountered an error: %v; continuing to do the next task.", err)
				continue
			}
			taskNum++
		}
	}
}

// sendNewTask sends a new task to the task manager contract.
func (avs *Avs) sendNewTask() error {
	avs.logger.Info("Avs sending new task")
	var taskPowerTotal sdkmath.LegacyDec
	var lastErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		taskPowerTotal, lastErr = avs.avsReader.GtAVSUSDValue(&bind.CallOpts{}, avs.avsAddress)

		if lastErr == nil && !taskPowerTotal.IsZero() && !taskPowerTotal.IsNegative() {
			break
		}

		if lastErr != nil {
			avs.logger.Error("Cannot get AVSUSDValue",
				"err", lastErr,
				"attempt", attempt,
				"max_attempts", maxRetries)
		} else {
			avs.logger.Info("AVS USD value is zero or negative",
				"value", taskPowerTotal,
				"attempt", attempt,
				"max_attempts", maxRetries)
		}

		if attempt == maxRetries {
			// panic("the voting power of AVS is zero or negative")
		}

		time.Sleep(retryDelay)
	}

	if taskPowerTotal.IsZero() || taskPowerTotal.IsNegative() {
		// panic("the voting power of AVS is zero or negative")
	}
	_, err := avs.avsWriter.CreateNewTask(
		context.Background(),
		GenerateRandomName(5),
		avs.taskResponsePeriod,
		avs.taskChallengePeriod,
		avs.thresholdPercentage,
		avs.taskStatisticalPeriod)

	if err != nil {
		avs.logger.Error("Avs failed to sendNewTask", "err", err)
		return err
	}
	return nil
}
func GenerateRandomName(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
