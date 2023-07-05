package main

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

const (
	// Pay-to-Witness-Public-Key-Hash  （bip84）
	P2WPKH_Path = `m/84'/0'/0'/0/`
	// Pay-to-Script-Hash (P2SH)
	P2SH_Path = `m/49'/0'/0'/0/`
	// ord
	P2TR_Path = `m/86'/0'/0'/0/`
	// Pay To Pubkey Hash
	P2PKH_Path = `m/44'/0'/0'/0/`

	// Pay-to-Witness-Script-Hash
	P2WSH_Path = `m/49'/0'/0'/0/`
)

// 按照路径派生子HD私钥
func deriveChildKey(masterKey *bip32.Key, path string) (*bip32.Key, error) {
	// 解析路径
	childPath, err := bip32.ParsePath(path)
	if err != nil {
		return nil, err
	}

	// 派生子HD私钥
	childKey := masterKey
	for _, index := range childPath {
		childKey, err = childKey.Child(index)
		if err != nil {
			return nil, err
		}
	}

	return childKey, nil
}
func main() {
	// entropy, _ := bip39.NewEntropy(256)
	// mnemonic, _ := bip39.NewMnemonic(entropy)
	// fmt.Println(mnemonic)
	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	mnemonic := "thrive member govern lake wealth alley theory divert roast screen poet maximum"
	seed := bip39.NewSeed(mnemonic, "")
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Show that the generated master node extended key is private.
	fmt.Println("Private Extended Key?:", masterKey.IsPrivate())
	// masterKey.ChainCode()
}
