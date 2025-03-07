package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"log"

	"LoadContract/store"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	contractAddr = "0x1399E0Fa83fBc9d1f34E16fe8603CF1348aA943B"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/9LURSvm6osXr98M_7j_AfY4fdhs2J9WL")
	if err != nil {
		log.Fatal(err)
	}
	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	//_ = storeContract
	//调用Version()
	version, err := storeContract.Version(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("version：", version)
	//调用Items()
	bytes, err := hexutil.Decode("0x0000000000000000000000000000000000000000000000000000000000001000")
	if err != nil {
		log.Fatal(err)
	}
	// 定义一个长度为 32 的字节数组
	var arg0 [32]byte
	// 将字节切片的数据复制到字节数组中
	/*
		arg0[:] 的含义
		arg0[:] 会把 arg0 这个数组转换为切片。在 Go 语言中，数组可以通过这种方式转换为切片，从而方便使用 copy 函数进行数据复制。
	*/
	copy(arg0[:], bytes)
	_item, err := storeContract.Items(&bind.CallOpts{}, arg0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("_item：", item)
	item := hexutil.Encode(_item[:]) //把[32]byte的_item转成切片[]byte
	fmt.Println("item：", item)
}
