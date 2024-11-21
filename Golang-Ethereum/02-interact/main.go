package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/2362250285044fa7a2c94e647daafb67"
var ganacheURL = "http://localhost:8545"

func main() {
	//last block number
	client, err := ethclient.DialContext(context.Background(), infuraURL)
	// client, err := ethclient.DialContext(context.Background(), ganacheURL)

	if err != nil {
		log.Fatalf("Error to create a ether client:%v", err)
	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block: %v", err)
	}
	fmt.Println(block.Number())

	addr := "0x0eCd45e023d49b250a7a9B41817c627E3Ef99eF9"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error to get the balance:%v", err)
	}
	fmt.Println("The balance:", balance)
	// 1 ether = 10^18 wei
	fBlance := new(big.Float)
	fBlance.SetString(balance.String())
	fmt.Println(fBlance)
	// 10*10*10*10*...18
	blanceEther := new(big.Float).Quo(fBlance, big.NewFloat(math.Pow10(18)))
	fmt.Println(blanceEther)

}
