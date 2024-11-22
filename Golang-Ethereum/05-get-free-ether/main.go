package main

import (
	"context"
	"fmt"
	"log"
	// "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	
)

var (
	url = "https://polygon-mainnet.infura.io/v3/2362250285044fa7a2c94e647daafb67"
)

func main() {
	// ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// _, err := ks.NewAccount("password")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = ks.NewAccount("password")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //b3fd25c6865aa6a09337b5c8392fd0518e478add
	// //26630da1a48014eed5f52cbab190210b245eccea

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	al := common.HexToAddress("b3fd25c6865aa6a09337b5c8392fd0518e478add")
	a2 := common.HexToAddress("26630da1a48014eed5f52cbab190210b245eccea")
	bl, err := client.BalanceAt(context.Background(), al, nil)
	if err != nil {
		log.Fatal(err)
	}
	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance 1:", bl)
	fmt.Println("Balance 2:", b2)

}