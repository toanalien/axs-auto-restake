package main

import (
	"axs-auto-stake/token"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"log"
	"math/big"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mnemonic := os.Getenv("MNEMONIC")
	rClient, _ := ethclient.Dial("https://api.roninchain.com/rpc")
	wClient, _ := ethclient.Dial("https://proxy.roninchain.com/free-gas-rpc")
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, true)
	if err != nil {
		log.Fatal(err)
	}

	_ = fmt.Sprintf("Try restake account: %s", account.Address.Hex())
	stakingManagerContract := common.HexToAddress("0x8bd81a19420bad681b7bfc20e703ebd8e253782d")
	axsStakingPoolContract := common.HexToAddress("0x05b0bb3c1c320b280501b86706c3551995bc8571")
	userAddress := common.HexToAddress(account.Address.Hex())

	// UserRewardInfo
	ins1, err := token.NewToken(stakingManagerContract, rClient)
	if err != nil {
		log.Fatal(err)
	}
	userRewardInfo, err := ins1.UserRewardInfo(&bind.CallOpts{}, axsStakingPoolContract, userAddress)
	if err != nil {
		log.Fatal(err)
	}
	lastClaimedBlock := userRewardInfo.LastClaimedBlock
	block, err := rClient.BlockByNumber(context.Background(), lastClaimedBlock)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		panic(err)
	}
	restakeTime := time.Unix(int64(block.Time()), 0).Add(time.Hour * time.Duration(24))
	localTime := time.Now().Local()

	if localTime.After(restakeTime) {
		// Restake
		nonce, err := rClient.PendingNonceAt(context.Background(), common.HexToAddress(account.Address.Hex()))
		ins2, err := token.NewToken(axsStakingPoolContract, wClient)
		if err != nil {
			log.Fatal(err)
		}
		key, err := wallet.PrivateKey(account)
		auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(2020))
		trans, err := ins2.RestakeRewards(&bind.TransactOpts{
			From:     userAddress,
			Nonce:    new(big.Int).SetUint64(nonce),
			Signer:   auth.Signer,
			GasPrice: big.NewInt(0),
			GasLimit: uint64(300000),
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("https://explorer.roninchain.com/tx/%s", trans.Hash().String()))
	}
}
