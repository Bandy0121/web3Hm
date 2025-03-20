package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/9LURSvm6osXr98M_7j_AfY4fdhs2J9WL")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("38edd948f94c292109f1ab6a8315a6d1c2934ee0f514047d41cb7bb25087ff4c")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) //类型断言
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(10000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)              // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x804c591679b1a49d9556AEF7BFE9E59f2265ffC0")
	var data []byte
	//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)  //NewTransaction已弃用
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//构建原始交易
	//rawTxBytes, err := signedTx.MarshalBinary()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//rawTxHex := hex.EncodeToString(rawTxBytes)
	//fmt.Printf(rawTxHex) // f86...772
	//
	//_rawTxBytes, err := hex.DecodeString(rawTxHex)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_tx := new(types.Transaction)
	//rlp.DecodeBytes(_rawTxBytes, &_tx)
	//err = client.SendTransaction(context.Background(), _tx)发送原始交易事务

	//广播交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	//0x7b045fce597436ebabbb0e29e339eac6642b586b29072f3bd4a6e34607ffea92
}
