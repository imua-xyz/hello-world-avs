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
	sdkecdsa "github.com/ExocoreNetwork/exocore-sdk/crypto/ecdsa"
	sdklogging "github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ethereum/go-ethereum"
	blscommon "github.com/prysmaticlabs/prysm/v4/crypto/bls/common"
	"math/big"
	"time"

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
	newTaskCreatedChan chan *avs.ContractavsserviceTaskCreated
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
		logger.Info("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
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
		newTaskCreatedChan: make(chan *avs.ContractavsserviceTaskCreated),
		avsAddr:            common.HexToAddress(c.AVSAddress),
	}

	if c.RegisterOperatorOnStartup {
		operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
			c.OperatorEcdsaPrivateKeyStorePath,
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

	firstHeight, err := o.ethClient.BlockNumber(context.Background())
	if err != nil {
		o.logger.Error("Cannot create AvsSubscriber", "err", err)
		return err
	}
	o.GetLog(int64(firstHeight))

	height := firstHeight
	fmt.Printf("Event firstHeight: %v\n", firstHeight)

	for {
		currentHeight, err := o.ethClient.BlockNumber(context.Background())
		fmt.Printf("Event currentHeight: %v\n", currentHeight)

		if err != nil {
			o.logger.Fatal(err.Error())
		}
		if currentHeight == height+1 {
			o.GetLog(int64(currentHeight))

			height = currentHeight
		}
		time.Sleep(2 * time.Second) // 等待一秒钟后再次查询
	}

	//case newTaskCreatedLog := <-o.newTaskCreatedChan:
	//	taskResponse := o.ProcessNewTaskCreatedLog(newTaskCreatedLog)
	//	sig, err := o.SignTaskResponse(taskResponse)
	//	if err != nil {
	//		continue
	//	}
	//	o.logger.Info("SignTaskResponse sig:", sig)
	//
	//	//go o.SendSignedTaskResponseToExocore(sig)

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
		contractAbi, _ := avs.ContractavsserviceMetaData.GetAbi()
		event := contractAbi.Events["TaskCreated"]
		for _, vLog := range logs {
			fmt.Println(vLog.BlockHash.Hex())
			fmt.Println(vLog.BlockNumber)
			fmt.Println(vLog.TxHash.Hex())
			data := vLog.Data
			fmt.Println(vLog.Topics)
			fmt.Println(event.ID)

			eventArgs, err := event.Inputs.Unpack(data)
			if err != nil {
				o.logger.Fatal(err.Error())
			}
			if eventArgs != nil {
				taskResponse := o.ProcessNewTaskCreatedLog(eventArgs)

				sig, err := o.SignTaskResponse(taskResponse)
				if err != nil {
					continue
				}
				fmt.Println(sig)
				//go o.SendSignedTaskResponseToExocore(sig)
			}
			//taskId := eventArgs[0].(*big.Int)
			//issuer := eventArgs[1].(common.Address)
			//name := eventArgs[2].(string)
			//taskResponsePeriod := eventArgs[3].(uint64)
			//taskChallengePeriod := eventArgs[4].(uint64)
			//thresholdPercentage := eventArgs[5].(uint64)
			//taskStatisticalPeriod := eventArgs[6].(uint64)
			//fmt.Println(data)
			//fmt.Println(eventArgs)
			//
			//fmt.Printf("Task ID: %v", taskId)
			//fmt.Printf("Issuer: %s", issuer.Hex())
			//fmt.Printf(name)
			//fmt.Printf("Task Response Period: %d", taskResponsePeriod)
			//fmt.Printf("Task Challenge Period: %d", taskChallengePeriod)
			//fmt.Printf("Threshold Percentage: %d", thresholdPercentage)
			//fmt.Printf("Task Statistical Period: %d", taskStatisticalPeriod)
		}
	}

}

// ProcessNewTaskCreatedLog Takes a NewTaskCreatedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the exocore as a task response.
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

func (o *Operator) SignTaskResponse(taskResponse *core.TaskResponse) (string, error) {

	taskResponseHash, err := core.GetTaskResponseDigestEncodeByjson(*taskResponse)
	if err != nil {
		o.logger.Error("Error SignTaskResponse with getting task response header hash. skipping task (this is not expected and should be investigated)", "err", err)
		return "", err
	}
	msgBytes := taskResponseHash[:]
	fmt.Println("ResHash:", hex.EncodeToString(msgBytes))

	sig := o.blsKeypair.Sign(msgBytes)

	sigStr := hex.EncodeToString(sig.Marshal())
	fmt.Println("sig:", sigStr)

	return sigStr, nil
}
