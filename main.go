package main

import (
	"context"
	"fmt"
	avs "github.com/ExocoreNetwork/exocore-avs/contracts/bindings/avs"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

const (
	wsURL       = "ws://localhost:8546"
	contractHex = "0x10Ed22D975453A5D4031440D51624552E4f204D5"
)

func main() {
	// 创建客户端连接
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}
	defer client.Close()

	// 创建事件过滤查询
	contractAddress := common.HexToAddress(contractHex)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	// 创建日志通道
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal("Subscribe failed:", err)
	}
	defer sub.Unsubscribe()

	fmt.Println("Starting event monitoring...")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("Subscription error:", err)
		case vLog := <-logs:
			// 解析日志
			event, err := parseEvent(vLog)
			if err != nil {
				log.Printf("Parse error: %v", err)
				continue
			}

			// 处理事件
			switch e := event.(type) {
			case *avs.ContracthelloWorldTaskCreated:
				fmt.Printf("New Task Created:\n"+
					"  TaskID: %v\n"+
					"  Issuer: %s\n"+
					"  Name: %s\n"+
					"  Number: %d\n"+
					"  Response Period: %d\n"+
					"  Challenge Period: %d\n"+
					"  Threshold: %d%%\n"+
					"  Statistical Period: %d\n",
					e.TaskId, e.Issuer.Hex(), e.Name, e.NumberToBeSquared,
					e.TaskResponsePeriod, e.TaskChallengePeriod,
					e.ThresholdPercentage, e.TaskStatisticalPeriod)

			case *avs.ContracthelloWorldTaskResolved:
				fmt.Printf("Task Resolved:\n"+
					"  TaskID: %d\n"+
					"  Address: %s\n",
					e.TaskId, e.TaskAddress.Hex())
			}
		}
	}
}

// 解析日志到具体事件
func parseEvent(vLog types.Log) (interface{}, error) {
	// 通过事件哈希判断事件类型
	contractABI, _ := avs.ContracthelloWorldMetaData.GetAbi()

	switch vLog.Topics[0] {
	case contractABI.Events["TaskCreated"].ID:
		return parseTaskCreated(vLog)
	case contractABI.Events["TaskResolved"].ID:
		return parseTaskResolved(vLog)
	default:
		return nil, fmt.Errorf("unknown event type")
	}
}

func parseTaskCreated(vLog types.Log) (*avs.ContracthelloWorldTaskCreated, error) {
	var event avs.ContracthelloWorldTaskCreated
	contractABI, _ := avs.ContracthelloWorldMetaData.GetAbi()

	// 由于都是非索引参数，直接解码Data字段
	err := contractABI.UnpackIntoInterface(&event, "TaskCreated", vLog.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack TaskCreated: %v", err)
	}

	return &event, nil
}

func parseTaskResolved(vLog types.Log) (*avs.ContracthelloWorldTaskResolved, error) {
	var event avs.ContracthelloWorldTaskResolved
	contractABI, _ := avs.ContracthelloWorldMetaData.GetAbi()

	err := contractABI.UnpackIntoInterface(&event, "TaskResolved", vLog.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack TaskResolved: %v", err)
	}

	return &event, nil
}
