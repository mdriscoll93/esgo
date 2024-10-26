package main

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

// GenerateMultisigAddress creates a 2-of-3 multisig address for Bitcoin escrow on Testnet
func GenerateMultisigAddress() (string, error) {
	// Generate three private keys (simulating buyer, seller, and arbitrator)
	privKey1, _ := secp256k1.GeneratePrivateKey()
	privKey2, _ := secp256k1.GeneratePrivateKey()
	privKey3, _ := secp256k1.GeneratePrivateKey()

	// Get the public keys from the private keys
	pubKey1 := privKey1.PubKey()
	pubKey2 := privKey2.PubKey()
	pubKey3 := privKey3.PubKey()

	// Convert public keys to AddressPubKey format for TestNet3
	addrPubKey1, err := btcutil.NewAddressPubKey(pubKey1.SerializeCompressed(), &chaincfg.TestNet3Params)
	if err != nil {
		return "", fmt.Errorf("failed to create address pub key 1: %w", err)
	}
	addrPubKey2, err := btcutil.NewAddressPubKey(pubKey2.SerializeCompressed(), &chaincfg.TestNet3Params)
	if err != nil {
		return "", fmt.Errorf("failed to create address pub key 2: %w", err)
	}
	addrPubKey3, err := btcutil.NewAddressPubKey(pubKey3.SerializeCompressed(), &chaincfg.TestNet3Params)
	if err != nil {
		return "", fmt.Errorf("failed to create address pub key 3: %w", err)
	}

	// Create the redeem script (2-of-3 multisig)
	redeemScript, err := txscript.MultiSigScript(
		[]*btcutil.AddressPubKey{addrPubKey1, addrPubKey2, addrPubKey3}, 2)
	if err != nil {
		return "", fmt.Errorf("failed to create redeem script: %w", err)
	}

	// Generate a P2SH address from the redeem script on Testnet
	scriptHash := btcutil.Hash160(redeemScript)
	address, err := btcutil.NewAddressScriptHashFromHash(scriptHash, &chaincfg.TestNet3Params)
	if err != nil {
		return "", fmt.Errorf("failed to create P2SH address: %w", err)
	}

	return address.EncodeAddress(), nil
}

func main() {
	fmt.Println("Creating a 2-of-3 multisig address for Bitcoin escrow on Testnet...")

	address, err := GenerateMultisigAddress()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("2-of-3 Multisig Testnet Bitcoin Address:", address)
}
