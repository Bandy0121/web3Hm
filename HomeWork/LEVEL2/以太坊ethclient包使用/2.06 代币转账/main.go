package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/9LURSvm6osXr98M_7j_AfY4fdhs2J9WL")
	if err != nil {
		log.Fatal(err)
	}

	txHash := common.HexToHash("0x7b045fce597436ebabbb0e29e339eac6642b586b29072f3bd4a6e34607ffea92")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			fmt.Println("Transaction not mined yet")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Transaction mined. Status: %d\n", receipt.Status)
	}

}
