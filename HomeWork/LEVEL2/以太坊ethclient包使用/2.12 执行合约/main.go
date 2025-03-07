package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
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

	contractAddr := common.HexToAddress("0x1399E0Fa83fBc9d1f34E16fe8603CF1348aA943B")
	// 从文件读取ABI内容也是可行的
	data, err := os.ReadFile("contractABI.json")
	if err != nil {
		log.Fatal(err)
	}
	var contractABI abi.ABI
	err = contractABI.UnmarshalJSON(data) //反序列化UnmarshalJSON
	if err != nil {
		log.Fatalf("Failed to unmarshal ABI: %v", err)
	}
	contract := bind.NewBoundContract(contractAddr, contractABI, client, client, client)

	//_ = contract
	//调用Version()
	var version []interface{}
	err = contract.Call(&bind.CallOpts{}, &version, "version")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("version：", version)

	//调用SetItem()
	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_save_key"))
	copy(value[:], []byte("demo_save_value11111"))

	//初始化交易opt实例
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	// 调用合约方法
	tx, err := contract.Transact(opt, "setItem", key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	//调用Items()
	var results []interface{}
	callOpt := &bind.CallOpts{Context: context.Background()}
	err = contract.Call(callOpt, &results, "items", key)
	if err != nil {
		log.Fatal(err)
	}
	if len(results) > 0 {
		if valueInContract, ok := results[0].([32]byte); ok {
			fmt.Println("valueInContract：", valueInContract)
			fmt.Println("is value saving in contract equals to origin value:", valueInContract == value)
			item := hexutil.Encode(valueInContract[:]) //把[32]byte的_item转成切片[]byte
			fmt.Println("item：", item)
		} else {
			log.Fatal("Unexpected result type")
		}
	} else {
		log.Fatal("No results returned")
	}
}
