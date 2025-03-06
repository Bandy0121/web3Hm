package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	//gasPrice, err := client.SuggestGasPrice(context.Background()) //获取平均燃气价格
	//if err != nil {
	//	log.Fatal(err)
	//}
	//gasPrice = big.NewInt(0).Add(gasPrice, big.NewInt(10000000000))
	gasPrice := big.NewInt(int64(60066570755))
	fmt.Println("gasPrice：", gasPrice) // 716826965
	toAddress := common.HexToAddress("0x804c591679b1a49d9556AEF7BFE9E59f2265ffC0")
	tokenAddress := common.HexToAddress("0x3058417327af0c01eb6eca30dd3902830f6d2f76")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	//预估燃气上限
	//gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	To:   &toAddress,
	//	Data: data,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(gasLimit) // 22946
	gasLimit := uint64(300000)
	//tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data) //NewTransaction已弃用
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &tokenAddress,
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

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	//0xf7326f448274aeb8745cca7dcb3357659a106d5edbaffa17172242ff436fb520
	//0x15a36895e2ba676617779abc371310671524ec3df3060c78c4ed9aaf096a977b
	//0xa944f6f80785b3c28494b3a77726796d50ddf722aa9fcd62cc47ecf21bea6116
	//0x481dc48c92f8f9bfe094aa567c14832c3a6b8ad27f9a5f72c5ea80b715446ce6
	//0x7b019a8d9c16dab05288682d5eb97e76f641c62ea81f410a2558089ef871f844
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
