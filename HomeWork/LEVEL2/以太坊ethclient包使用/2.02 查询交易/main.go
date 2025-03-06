package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/9LURSvm6osXr98M_7j_AfY4fdhs2J9WL")
	if err != nil {
		log.Fatal(err)
	}

	//chainID, err := client.ChainID(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//blockNumber := big.NewInt(5671744)
	//block, err := client.BlockByNumber(context.Background(), blockNumber)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//0xf7326f448274aeb8745cca7dcb3357659a106d5edbaffa17172242ff436fb520
	//0x15a36895e2ba676617779abc371310671524ec3df3060c78c4ed9aaf096a977b
	txHash := common.HexToHash("0xa944f6f80785b3c28494b3a77726796d50ddf722aa9fcd62cc47ecf21bea6116")
	receipt, err := client.TransactionReceipt(context.Background(), txHash) //tx.Hash()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt.Status) // 1
	fmt.Println(receipt.Logs)

	//for _, tx := range block.Transactions() {
	//	fmt.Println(tx.Hash().Hex())        // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
	//	fmt.Println(tx.Value().String())    // 100000000000000000
	//	fmt.Println(tx.Gas())               // 21000
	//	fmt.Println(tx.GasPrice().Uint64()) // 100000000000
	//	fmt.Println(tx.Nonce())             // 245132
	//	fmt.Println(tx.Data())              // []
	//	fmt.Println(tx.To().Hex())          // 0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587
	//
	//	if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
	//		fmt.Println("sender", sender.Hex()) // 0x2CdA41645F2dBffB852a605E92B185501801FC28
	//	} else {
	//		log.Fatal(err)
	//	}
	//
	//	//txHash := common.HexToHash("0xa944f6f80785b3c28494b3a77726796d50ddf722aa9fcd62cc47ecf21bea6116")
	//	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash()) //txHash
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(receipt.Status) // 1
	//	fmt.Println(receipt.Logs)   // []
	//	break
	//}

	//blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	//count, err := client.TransactionCount(context.Background(), blockHash)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for idx := uint(0); idx < count; idx++ {
	//	tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
	//	break
	//}
	//
	//txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	//tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(isPending)
	//fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5.Println(isPending)       // false
}
