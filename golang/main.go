package main

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
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

func main() {
	// entropy, _ := bip39.NewEntropy(256)
	// mnemonic, _ := bip39.NewMnemonic(entropy)
	// fmt.Println(mnemonic)
	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	mnemonic := "thrive member govern lake wealth alley theory divert roast screen poet maximum"
	seed := bip39.NewSeed(mnemonic, "")

	// p2wpkh
	p2wpkhKey, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)

	p2wpkhFirst, _ := p2wpkhKey.Derive(hdkeychain.HardenedKeyStart + 84)
	p2wpkhSecond, _ := p2wpkhFirst.Derive(hdkeychain.HardenedKeyStart + 0)
	p2wpkhThird, _ := p2wpkhSecond.Derive(hdkeychain.HardenedKeyStart + 0)
	p2wpkhPre, _ := p2wpkhThird.Derive(0)
	p2wpkhFull, _ := p2wpkhPre.Derive(0)
	p2wpkhPub, _ := p2wpkhFull.ECPubKey()

	p2wpkhPubHex := fmt.Sprintf("%x", p2wpkhPub.SerializeCompressed())
	fmt.Println("p2wpkh格式生成的公钥为：", p2wpkhPubHex)
	p2wpkhStingAddress, _ := p2wpkhFull.Address(&chaincfg.MainNetParams)
	p2wpkhAddress, _ := btcutil.NewAddressWitnessPubKeyHash(p2wpkhStingAddress.ScriptAddress(), &chaincfg.MainNetParams)
	fmt.Println("p2wpkh格式生成的地址为：", p2wpkhAddress.EncodeAddress())
	fmt.Println("---")

	// p2sh-p2wpkh
	p2shP2wpkhKey, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	p2shP2wpkhFirst, _ := p2shP2wpkhKey.Derive(hdkeychain.HardenedKeyStart + 49)
	p2shP2wpkhSecond, _ := p2shP2wpkhFirst.Derive(hdkeychain.HardenedKeyStart + 0)
	p2shP2wpkhThird, _ := p2shP2wpkhSecond.Derive(hdkeychain.HardenedKeyStart + 0)
	p2shP2wpkhPre, _ := p2shP2wpkhThird.Derive(0)
	p2shP2wpkhFull, _ := p2shP2wpkhPre.Derive(0)
	p2shP2wpkhPub, _ := p2shP2wpkhFull.ECPubKey()
	p2shP2wpkhPubHex := fmt.Sprintf("%x", p2shP2wpkhPub.SerializeCompressed())
	fmt.Println("p2sh-p2wpkh格式生成的公钥为：", p2shP2wpkhPubHex)
	p2shP2wpkhStringAddress, _ := p2shP2wpkhFull.Address(&chaincfg.MainNetParams)
	p2shP2wpkhPubKeyHash, _ := btcutil.NewAddressWitnessPubKeyHash(p2shP2wpkhStringAddress.ScriptAddress(), &chaincfg.MainNetParams)
	p2shP2wpkhScript, _ := txscript.PayToAddrScript(p2shP2wpkhPubKeyHash)
	p2shP2wpkhResult, _ := btcutil.NewAddressScriptHash(p2shP2wpkhScript, &chaincfg.MainNetParams)
	fmt.Println("p2sh-p2wpkh格式生成的地址为：", p2shP2wpkhResult.EncodeAddress())
	fmt.Println("---")

	// p2tr
	p2trKey, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	p2trFirst, _ := p2trKey.Derive(hdkeychain.HardenedKeyStart + 86)
	p2trSecond, _ := p2trFirst.Derive(hdkeychain.HardenedKeyStart + 0)
	p2trThird, _ := p2trSecond.Derive(hdkeychain.HardenedKeyStart + 0)
	p2trPre, _ := p2trThird.Derive(0)
	p2trFull, _ := p2trPre.Derive(0)
	p2trPub, _ := p2trFull.ECPubKey()
	p2trPubHex := fmt.Sprintf("%x", p2trPub.SerializeCompressed())
	fmt.Println("p2tr格式生成的公钥为：", p2trPubHex)
	p2trTapKey := txscript.ComputeTaprootKeyNoScript(p2trPub)
	p2trBytes := p2trTapKey.SerializeCompressed()[1:33]
	p2trResult, _ := btcutil.NewAddressTaproot(p2trBytes, &chaincfg.MainNetParams)
	fmt.Println("p2tr格式生成的地址为：", p2trResult.EncodeAddress())
	fmt.Println("---")

	// p2pkh
	p2pkhKey, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)

	p2pkhFirst, _ := p2pkhKey.Derive(hdkeychain.HardenedKeyStart + 44)
	p2pkhSecond, _ := p2pkhFirst.Derive(hdkeychain.HardenedKeyStart + 0)
	p2pkhThird, _ := p2pkhSecond.Derive(hdkeychain.HardenedKeyStart + 0)
	p2pkhPre, _ := p2pkhThird.Derive(0)
	p2pkhFull, _ := p2pkhPre.Derive(0)
	p2pkhPub, _ := p2pkhFull.ECPubKey()
	p2pkhPubHex := fmt.Sprintf("%x", p2pkhPub.SerializeCompressed())
	fmt.Println("p2pkh格式生成的公钥为：", p2pkhPubHex)

	p2pkhAddress, _ := p2pkhFull.Address(&chaincfg.MainNetParams)
	fmt.Println("p2pkh格式生成的地址为：", p2pkhAddress.String())

}
