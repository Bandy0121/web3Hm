package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

//var StoreABI = `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

func main() {
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/9LURSvm6osXr98M_7j_AfY4fdhs2J9WL")
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress("0x1399E0Fa83fBc9d1f34E16fe8603CF1348aA943B")
	// 从文件读取ABI内容也是可行的
	data, err := os.ReadFile("../contractABI.json")
	if err != nil {
		log.Fatal(err)
	}
	var contractAbi abi.ABI
	err = contractAbi.UnmarshalJSON(data) //反序列化UnmarshalJSON
	if err != nil {
		log.Fatalf("Failed to unmarshal ABI: %v", err)
	}
	//contractAbi, err := abi.JSON(strings.NewReader(string(StoreABI)))
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 定义要订阅的事件签名
	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	eventHash := crypto.Keccak256Hash(eventSignature)

	// 创建过滤器查询
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{eventHash}},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog.BlockHash.Hex())
			fmt.Println(vLog.BlockNumber)
			fmt.Println(vLog.TxHash.Hex())
			event := struct {
				Key   [32]byte
				Value [32]byte
			}{}
			err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(common.Bytes2Hex(event.Key[:]))
			fmt.Println(common.Bytes2Hex(event.Value[:]))
			var topics []string
			for i := range vLog.Topics {
				topics = append(topics, vLog.Topics[i].Hex())
			}
			fmt.Println("topics[0]=", topics[0])
			if len(topics) > 1 {
				fmt.Println("index topic:", topics[1:])
			}
		}
	}
}
