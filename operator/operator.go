package operator

import (
	"context"
	"encoding/hex"
	cstaskmanager "github.com/ExocoreNetwork/exocore-avs/contracts/bindings/avs"
	"github.com/ExocoreNetwork/exocore-avs/core"
	chain "github.com/ExocoreNetwork/exocore-avs/core/chainio"
	"github.com/ExocoreNetwork/exocore-avs/core/chainio/eth"
	"github.com/ExocoreNetwork/exocore-avs/types"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	"github.com/ExocoreNetwork/exocore-sdk/crypto/bls"
	sdkecdsa "github.com/ExocoreNetwork/exocore-sdk/crypto/ecdsa"
	sdklogging "github.com/ExocoreNetwork/exocore-sdk/logging"
	blscommon "github.com/prysmaticlabs/prysm/v4/crypto/bls/common"

	"github.com/ExocoreNetwork/exocore-sdk/nodeapi"
	"github.com/ExocoreNetwork/exocore-sdk/signerv2"
	"github.com/ethereum/go-ethereum/common"
	"os"
)

const AvsName = "hello-world-avs-demo"
const SemVer = "0.0.1"

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
	newTaskCreatedChan chan *cstaskmanager.ContractavsserviceTaskCreated
	// needed when opting in to avs (allow this service manager contract to slash operator)
	avsAddr common.Address
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
		logger.Errorf("Cannot create http ethclient", "err", err)
		return nil, err
	}
	ethWsClient, err = eth.NewClient(c.EthWsUrl)
	if err != nil {
		logger.Errorf("Cannot create ws ethclient", "err", err)
		return nil, err
	}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(c.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		logger.Errorf("Cannot parse bls private key", "err", err)
		return nil, err
	}
	// TODO(samlaf): should we add the chainId to the config instead?
	// this way we can prevent creating a signer that signs on mainnet by mistake
	// if the config says chainId=5, then we can only create a goerli signer
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

	txMgr := txmgr.NewSimpleTxManager(ethRpcClient, logger, signerV2, common.HexToAddress(c.OperatorAddress))

	avsReader, _ := chain.BuildExoChainReader(
		common.HexToAddress(c.AVSAddress),
		ethRpcClient,
		logger)

	avsWriter, _ := chain.BuildELChainWriter(
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
		newTaskCreatedChan: make(chan *cstaskmanager.ContractavsserviceTaskCreated),
		avsAddr:            common.HexToAddress(c.AVSAddress),
	}

	if c.RegisterOperatorOnStartup {
		operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
			c.EcdsaPrivateKeyStorePath,
			ecdsaKeyPassword,
		)
		if err != nil {
			return nil, err
		}
		operator.registerOperatorOnStartup(operatorEcdsaPrivateKey)
	}

	logger.Info("Operator info",
		"operatorAddr", c.OperatorAddress,
		"operatorKey", operator.blsKeypair.PublicKey().Marshal(),
	)

	return operator, nil

}

func (o *Operator) Start(ctx context.Context) error {
	//operatorIsRegistered, err := o.avsReader.IsOperatorRegistered(&bind.CallOpts{}, o.operatorAddr)
	//if err != nil {
	//	o.logger.Error("Error checking if operator is registered", "err", err)
	//	return err
	//}
	//if !operatorIsRegistered {
	//	// We bubble the error all the way up instead of using logger.Fatal because logger.Fatal prints a huge stack trace
	//	// that hides the actual error message. This error msg is more explicit and doesn't require showing a stack trace to the user.
	//	return fmt.Errorf("operator is not registered. Registering operator using the operator-cli before starting operator")
	//}

	o.logger.Infof("Starting operator.")

	if o.config.EnableNodeApi {
		o.nodeApi.Start()
	}

	sub := o.avsSubscriber.SubscribeToNewTasks(o.newTaskCreatedChan)
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-sub.Err():
			o.logger.Error("Error in websocket subscription", "err", err)
			// TODO(samlaf): write unit tests to check if this fixed the issues we were seeing
			sub.Unsubscribe()
			// TODO(samlaf): wrap this call with increase in avs-node-spec metric
			sub = o.avsSubscriber.SubscribeToNewTasks(o.newTaskCreatedChan)
		case newTaskCreatedLog := <-o.newTaskCreatedChan:
			taskResponse := o.ProcessNewTaskCreatedLog(newTaskCreatedLog)
			sig, err := o.SignTaskResponse(taskResponse)
			if err != nil {
				continue
			}
			o.logger.Info("SignTaskResponse sig:", sig)

			//go o.SendSignedTaskResponseToExocore(sig)
		}
	}
}

// ProcessNewTaskCreatedLog Takes a NewTaskCreatedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the contract as a task response.
func (o *Operator) ProcessNewTaskCreatedLog(newTaskCreatedLog *cstaskmanager.ContractavsserviceTaskCreated) *core.TaskResponse {
	o.logger.Debug("Received new task", "task", newTaskCreatedLog)
	o.logger.Info("Received new task",
		"numberToBeSquared", newTaskCreatedLog.TaskId,
		"taskIndex", newTaskCreatedLog.Name,
	)
	taskResponse := &core.TaskResponse{
		TaskID: newTaskCreatedLog.TaskId,
		Msg:    newTaskCreatedLog.Name,
	}
	return taskResponse
}

func (o *Operator) SignTaskResponse(taskResponse *core.TaskResponse) (string, error) {
	taskResponseHash, err := core.GetTaskResponseDigest(taskResponse)
	if err != nil {
		o.logger.Error("Error getting task response header hash. skipping task (this is not expected and should be investigated)", "err", err)
		return "", err
	}
	sig := o.blsKeypair.Sign(taskResponseHash[:])
	sigStr := hex.EncodeToString(sig.Marshal())
	return sigStr, nil
}
