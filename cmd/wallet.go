package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"

	"github.com/anonnice1999/questx-script/internal/wallet"
)

func startWallet(cCtx *cli.Context) error {
	secret := getEnv("BLOCKCHAIN_SECRET_KEY", "eth_super_super_secret_key_should_be_32_bytes")
	nonce := cCtx.String("nonce")

	privateKey, err := wallet.GeneratePrivateKey([]byte(secret), []byte(nonce))
	if err != nil {
		panic(err)
	}

	fmt.Printf("private key: %x\n", privateKey.D.Bytes())
	fmt.Println("public key: ", crypto.PubkeyToAddress(privateKey.PublicKey))
	return nil
}
