package main

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethclient "github.com/ethereum/go-ethereum/ethClient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {

	contractAddress := "0x71cbeC8ab8686bCA859DF356860a62917B92fc9C"
	rpcURL := "http://127.0.0.1:8545"
	ganacheClient := &ethclient.Client{}

	if ganacheRPCClient, err := rpc.Dial(rpcURL); err == nil {
		ganacheClient = ethclient.NewClient(ganacheRPCClient)
	} else {
		panic(err)
	}

	storageInstance, err := NewStorage(common.HexToAddress(contractAddress), ganacheClient)
	if err != nil {
		panic(err)
	}
	val, err := storageInstance.GetVal(nil)
	if err != nil {
		panic(err)
	}
	log.Println("value is - ", val)

	storageInstance.SetVal(nil, big.NewInt(5))
	val, err = storageInstance.GetVal(nil)
	log.Println("value is - ", val)

}
