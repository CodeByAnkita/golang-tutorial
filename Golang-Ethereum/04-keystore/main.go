package main

import (
	"fmt"

	"log"

	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "password" // if this password change then you not get wallet folder from al keys
	// // for that you need to first run commanded code for new wallet
	// a, err := key.NewAccount(password)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(a.Address)

	// // Read the keystore file
	b, err := ioutil.ReadFile("./wallet/UTC--2024-11-21T10-31-09.722872000Z--c8830858abe860fecb540de206b7d6032ae37516")
	//c8830858abe860fecb540de206b7d6032ae37516
	if err != nil {
		log.Fatal("Error reading the keystore file:", err)
	}

	// Decrypt the keystore to retrieve the private key
	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal("Error decrypting the keystore file:", err)
	}

	// Extract the private key
	privateKeyData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Private Key:", hexutil.Encode(privateKeyData)) // Hex-encoded private key

	// Extract the public key
	publicKeyData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Public Key:", hexutil.Encode(publicKeyData)) // Hex-encoded public key

	// Extract the address
	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	fmt.Println("Address:", address)
}
