package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mnemonic := os.Getenv("MNEMONIC")
	axsStakingPoolAddr := "0x05b0bb3c1c320b280501b86706c3551995bc8571"
	client, err := ethclient.Dial("https://proxy.roninchain.com/free-gas-rpc")
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
	value := big.NewInt(0)
	toAddress := common.HexToAddress(axsStakingPoolAddr)
	gasLimit := uint64(300000)
	gasPrice := big.NewInt(0)
	var data []byte
	transferFnSignature := []byte("restakeRewards()")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	data = append(data, methodID...)

	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(account.Address.Hex()))
	if err != nil {
		log.Fatal(err)
	}
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	chainID := big.NewInt(2020)
	signedTx, err := wallet.SignTxEIP155(account, tx, chainID)
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
}
