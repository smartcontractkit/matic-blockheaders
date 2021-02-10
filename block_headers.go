package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	client, err := ethclient.Dial("https://rpc-mainnet.maticvigil.com/")
	if err != nil {
		log.Fatal(err)
	}

	latestBlockHeader, err := client.HeaderByNumber(context.Background(), nil) // most recent block
	printBlockRLP(latestBlockHeader)

	noTestCases := 5
	blockNum := big.NewInt(10000467)
	for i := 0; i < noTestCases; i++ {
		if err != nil {
			log.Fatal((err))
		}
		header, err := client.HeaderByNumber(context.Background(), blockNum)
		if err != nil {
			log.Fatal(err)
		}
		printBlockRLP(header)
		blockNum.Add(blockNum, big.NewInt(1))
	}
}

func printBlockRLP(header *types.Header) {
	data, err := rlp.EncodeToBytes(header)
	if err != nil {
		log.Fatal(err)
	}
	output := hex.EncodeToString(data)
	fmt.Printf("Block Number: %v\n", header.Number)
	fmt.Printf("Parent hash: %v\n", header.ParentHash.Hex())
	fmt.Printf("0x%s\n\n", output)
}
