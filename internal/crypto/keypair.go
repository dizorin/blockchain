package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

type KeyPair struct {
	PrivateKey crypto.PrivateKey
	PublicKey  crypto.PublicKey
}

func CreateKeyPair() *KeyPair {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return &KeyPair{key, key.Public()}
}

func (k KeyPair) String() string {
	privateKey := k.PrivateKey.(*ecdsa.PrivateKey)
	return fmt.Sprintf("[PublicKey=%x:%x,PrivateKey=%x]", privateKey.X, privateKey.Y, privateKey.D)
}
