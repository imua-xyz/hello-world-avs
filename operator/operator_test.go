package operator_test

import (
	"context"
	"fmt"
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
		time.Sleep(1 * time.Second) // 等待一秒钟后再次查询
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
		for _, log := range logs {
			fmt.Printf("Event Received: %v\n", log)
			// 在这里处理事件日志
		}
	}

}
