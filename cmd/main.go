package main

import (
	"fmt"
	"github.com/dizorin/blockchain/internal/crypto"
)

func main() {
	text := []byte("TEXT")
	fmt.Println("text =", string(text))

	pair := crypto.CreateKeyPair()
	fmt.Println("pair =", pair)

	sign, err := crypto.SignData(pair.PrivateKey, text)
	if err != nil {
		panic(err)
	}
	fmt.Println("sign =", sign)

	fmt.Println("Verify =", crypto.VerifySign(sign, []byte("TEXT"), pair.PublicKey))
}
