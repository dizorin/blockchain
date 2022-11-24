package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
)

type Sign struct {
	r []byte
	s []byte
}

func SignData(privateKey crypto.PrivateKey, mes []byte) (Sign, error) {
	r, s, err := ecdsa.Sign(rand.Reader, privateKey.(*ecdsa.PrivateKey), mes)
	if err != nil {
		return Sign{}, err
	}
	return Sign{r.Bytes(), s.Bytes()}, nil
}

func VerifySign(sign Sign, mes []byte, publicKey crypto.PublicKey) bool {
	return ecdsa.Verify(publicKey.(*ecdsa.PublicKey), mes, new(big.Int).SetBytes(sign.r), new(big.Int).SetBytes(sign.s))
}

func (s Sign) String() string {
	return fmt.Sprintf("[r=%x,s=%x]", s.r, s.s)
}
