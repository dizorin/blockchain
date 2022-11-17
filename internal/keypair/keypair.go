package keypair

import "crypto"

type KeyPair struct {
	PrivateKey crypto.PrivateKey
	PublicKey  crypto.PublicKey
}

func createKeyPair() *KeyPair {
	return &KeyPair{}
}

func (*KeyPair) Print() string {
	return ""
}
