package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/2362250285044fa7a2c94e647daafb67"
var ganacheURL = "http://localhost:8545"

func main() {
	//last block number
	// client, err := ethclient.DialContext(context.Background(), infuraURL)
	client, err := ethclient.DialContext(context.Background(), ganacheURL)

	if err != nil {
		log.Fatalf("Error to create a ether client:%v", err)
	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block: %v", err)
	}
	fmt.Println(block.Number())
}
