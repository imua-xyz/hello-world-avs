package operator_test

import (
	"context"
	"fmt"
	avs "github.com/ExocoreNetwork/exocore-avs/contracts/bindings/avs"
	"github.com/ExocoreNetwork/exocore-avs/core"
	"github.com/ExocoreNetwork/exocore-avs/core/chainio/eth"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"testing"
	"time"
)

func TestAbi(t *testing.T) {

	task := core.TaskResponse{
		TaskID:        10,
		NumberSquared: 56169,
	}

	packed, err := core.Args.Pack(&task)
	if err != nil {
		t.Errorf("Error packing task: %v", err)
		return
	} else {
		t.Logf("ABI encoded: %s", hexutil.Encode(packed))
	}

	args := make(map[string]interface{})

	err = core.Args.UnpackIntoMap(args, packed)
	result, _ := core.Args.Unpack(packed)
	t.Logf("Unpacked: %v", result[0])
	hash := crypto.Keccak256Hash(packed)
	t.Logf("Hash: %s", hash.String())

	key := args["TaskResponse"]
	t.Logf("Key: %v", key)

}

// doc：https://docs.infura.io/api/networks/ethereum/json-rpc-methods/eth_getlogs
func TestEth_getlogs(t *testing.T) {
	ethRpcClient, err := eth.NewClient("http://localhost:8545")
	if err != nil {
		log.Fatal("Cannot create http ethclient", "err", err)
	}
	// Contract address and ABI
	contractAddress := common.HexToAddress("0x10Ed22D975453A5D4031440D51624552E4f204D5")
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	firstHeight, err := ethRpcClient.BlockNumber(context.Background())

	GetLog(ethRpcClient, contractAddress, int64(firstHeight))

	height := firstHeight
	fmt.Printf("Event firstHeight: %v\n", firstHeight)

	for {
		currentHeight, err := ethRpcClient.BlockNumber(context.Background())
		fmt.Printf("Event currentHeight: %v\n", currentHeight)

		if err != nil {
			log.Fatal(err)
		}
		if currentHeight == height+1 {
			GetLog(ethRpcClient, contractAddress, int64(currentHeight))

			height = currentHeight
		}
		time.Sleep(2 * time.Second)
	}
}

func GetLog(client eth.EthClient, address common.Address, height int64) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
		FromBlock: big.NewInt(height),
		ToBlock:   big.NewInt(height),
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	if logs != nil {
		contractAbi, _ := avs.ContracthelloWorldMetaData.GetAbi()
		event := contractAbi.Events["TaskCreated"]

		for _, vLog := range logs {
			fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
			fmt.Println(vLog.BlockNumber)     // 2394201
			fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6
			data := vLog.Data
			fmt.Println(vLog.Topics)
			fmt.Println(event.ID)

			eventArgs, err := event.Inputs.Unpack(data)

			log.Println("parse logs",
				"data", data,
				"height", height,
				"event", event.Inputs)
			if err != nil {
				log.Fatal(err)
			}

			// taskId := eventArgs[0].(*big.Int)
			issuer := eventArgs[1].(common.Address)
			name := eventArgs[2].(string)
			taskResponsePeriod := eventArgs[3].(uint64)
			taskChallengePeriod := eventArgs[4].(uint64)
			thresholdPercentage := eventArgs[5].(uint64)
			taskStatisticalPeriod := eventArgs[6].(uint64)
			fmt.Println(data)
			fmt.Println(eventArgs)

			// log.Printf("Task ID: %v", taskId)
			fmt.Printf("Issuer: %s", issuer.Hex())
			fmt.Printf(name)
			fmt.Printf("Task Response Period: %d", taskResponsePeriod)
			fmt.Printf("Task Challenge Period: %d", taskChallengePeriod)
			fmt.Printf("Threshold Percentage: %d", thresholdPercentage)
			fmt.Printf("Task Statistical Period: %d", taskStatisticalPeriod)
		}
	}

}
