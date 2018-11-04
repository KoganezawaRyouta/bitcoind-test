package interfaces

import (
	"flag"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
)

var testnet = flag.Bool("testnet", false, "operate on the testnet Bitcoin network")

// By default (without -testnet), use mainnet.
var chainParams = &chaincfg.MainNetParams

func GetInfo() error {
	flag.Parse()
	// Modify active network parameters if operating on testnet.
	chainParams = &chaincfg.TestNet3Params

	seed, err := hdkeychain.GenerateSeed(hdkeychain.RecommendedSeedLen)
	log.Printf(" seed : %v\n", seed)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Generate a new master node using the seed.
	master, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		panic(err)
	}
	// m/49'
	purpose, err := master.Child(49 + hdkeychain.HardenedKeyStart)
	if err != nil {
		panic(err)
	}
	// m/49'/1'
	coinType, err := purpose.Child(1 + hdkeychain.HardenedKeyStart)
	if err != nil {
		panic(err)
	}

	// m/49'/1'/0'
	account, err := coinType.Child(1 + hdkeychain.HardenedKeyStart)
	if err != nil {
		return err
	}
	// m/49'/1'/0'/0
	change, err := account.Child(0)
	if err != nil {
		return err
	}

	// m/49'/1'/0'/0/0
	addressIndex, err := change.Child(0)
	if err != nil {
		return err
	}

	extendedKey, err := addressIndex.Neuter()
	if err != nil {
		fmt.Println(err)
		return err
	}
	privateKey, err := extendedKey.ECPrivKey()
	if err != nil {
		return err
	}
	ecPubKey, err := extendedKey.ECPubKey()
	if err != nil {
		fmt.Println(err)
		return err
	}
	/// ============================================================================P2SH
	// payToWitnessPubKeyHashScript
	keyHash := btcutil.Hash160(ecPubKey.SerializeCompressed())
	ad, err := btcutil.NewAddressScriptHash(keyHash, chainParams)
	if err != nil {
		return err
	}

	wif, err := btcutil.NewWIF(privateKey, &chaincfg.TestNet3Params, true)
	if err != nil {
		return err
	}

	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:18332",
		User:         "test-net-user-rk2826",
		Pass:         "PT-78ba4BaBjhwQMQR2xh6VIAAYlNdxYHbj6NaPZUnc=",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	client.ImportPrivKeyLabelAsync(wif, "hello").Receive()
	// 公開鍵
	log.Printf(" Public key : %s\n")

	// 公開鍵
	log.Printf(" Public key : %s\n", ecPubKey.ToECDSA())
	// 秘密鍵
	log.Printf(" Private key : %v\n", addressIndex.String())
	// アドレス
	log.Printf(" EncodeAddress(P2SH) base58: %v\n", ad.String())
	//// 公開鍵ハッシュ
	//log.Printf(" PrivKey hash: %v\n", ecPubKey.SerializeUncompressed())
	//log.Printf(" PrivKey hash(comp) : %v\n", ecPubKey.SerializeCompressed())
	//
	//// 1. 署名者のアカウントの公開鍵よりAddressPubKeyを作成する
	//addressPubKey1, err := btcutil.NewAddressPubKey(ecPubKey.SerializeCompressed(), &chaincfg.MainNetParams)
	//if err != nil {
	//	return err
	//}
	//addressPubKey2, err := btcutil.NewAddressPubKey(ecPubKey2.SerializeCompressed(), &chaincfg.MainNetParams)
	//if err != nil {
	//	return err
	//}
	//addressPubKeys := []*btcutil.AddressPubKey{addressPubKey1, addressPubKey2}
	//// 　2. AddressPubKeyを使ってredeemScriptを作成する
	//redeemScript, err := txscript.MultiSigScript(addressPubKeys, len(addressPubKeys))
	//if err != nil {
	//	return err
	//}
	//// 3. redeemScriptよりp2sh adddressを作成する
	//ad, err := btcutil.NewAddressScriptHash(redeemScript, chainParams)
	//if err != nil {
	//	return err
	//}
	//addr := ad.EncodeAddress()

	//log.Printf(" ECPrivKey : %v\n", ecPrivKey.Serialize())
	//log.Printf(" ECPrivKey X: %v\n", ecPrivKey.X)
	//log.Printf(" ECPrivKey Y: %v\n", ecPrivKey.Y)
	//
	//log.Printf(" ECPubKey : %v\n", ecPubKey)
	//log.Printf(" ECPubKey X : %v\n", ecPubKey.X)
	//log.Printf(" ECPubKey Y : %v\n", ecPubKey.Y)

	//log.Printf(" EncodeAddress(P2SH) base58: %v\n", ad.String())
	//log.Printf(" EncodeAddress(P2SH) base58: %v\n", string(addr))
	//log.Printf(" Hash160(P2SH) : %v\n", ad.Hash160())

	///// ============================================================================P2SH
	//ad, err = btcutil.NewAddressScriptHash(ecPubKey.SerializeUncompressed(), chainParams)
	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	//log.Printf(" EncodeAddress(P2SH) : %v\n", ad.String())
	//log.Printf(" encodeAddress(P2SH) base58 : %v\n", ad.EncodeAddress())
	//
	///// ============================================================================P2SH
	//he := hex.EncodeToString(ecPubKey.SerializeUncompressed())
	//ret, err := hex.DecodeString(he)
	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	//ad, err = btcutil.NewAddressScriptHash(ret, chainParams)
	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	//log.Printf("ads EncodeAddress(P2SH) : %v\n", ad.String())
	//log.Printf("ads encodeAddress(P2SH) base58 : %v\n", ad.EncodeAddress())
	//
	///// ============================================================================P2WSH!
	//he = hex.EncodeToString(ecPubKey.SerializeUncompressed())
	//ret, err = hex.DecodeString(he)
	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	//ads, err := btcutil.NewAddressWitnessScriptHash(ret, chainParams)
	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	//log.Printf("ads EncodeAddress(P2WSH) : %v\n", ads.String())
	//log.Printf("ads encodeAddress(P2WSH) base58 : %v\n", ads.EncodeAddress())
	return nil
}
