package main

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

func main() {
	fmt.Println("Setting up a Bitcoin escrow service in Go")

	// example - confirm dependencies
	// example - generate new btc private key for testing
	privKey, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}
	fmt.Println("Private Key Generated:", privKey)

	// Generate a simple Bitcoin address (test code)
	addr, err := btcutil.NewAddressPubKey(privKey.PubKey().SerializeCompressed(), &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println("Error generating address:", err)
		return
	}
	fmt.Println("Bitcoin Address:", addr.EncodeAddress())

}
