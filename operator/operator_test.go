package operator_test

import (
	"context"
	"fmt"
	avs "github.com/ExocoreNetwork/exocore-avs/contracts/bindings/avs"
	"github.com/ExocoreNetwork/exocore-avs/core/chainio/eth"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"testing"
	"time"
)

// doc：https://docs.infura.io/api/networks/ethereum/json-rpc-methods/eth_getlogs
func TestEth_getlogs(t *testing.T) {
	ethRpcClient, err := eth.NewClient("http://localhost:8545")
	if err != nil {
		log.Fatal("Cannot create http ethclient", "err", err)
	}
	// 合约地址和 ABI
	contractAddress := common.HexToAddress("0xDF907c29719154eb9872f021d21CAE6E5025d7aB")
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
		time.Sleep(2 * time.Second) // 等待一秒钟后再次查询
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
		contractAbi, _ := avs.ContractavsserviceMetaData.GetAbi()
		event := contractAbi.Events["TaskCreated"]

		for _, vLog := range logs {
			fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
			fmt.Println(vLog.BlockNumber)     // 2394201
			fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6
			data := vLog.Data
			fmt.Println(vLog.Topics)
			fmt.Println(event.ID)

			eventArgs, err := event.Inputs.Unpack(data)
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

			//log.Printf("Task ID: %v", taskId)
			fmt.Printf("Issuer: %s", issuer.Hex())
			fmt.Printf(name)
			fmt.Printf("Task Response Period: %d", taskResponsePeriod)
			fmt.Printf("Task Challenge Period: %d", taskChallengePeriod)
			fmt.Printf("Threshold Percentage: %d", thresholdPercentage)
			fmt.Printf("Task Statistical Period: %d", taskStatisticalPeriod)
		}
	}

}
