package crypto

import "crypto/sha1"

func ToSHA1(mes []byte) []byte {
	return sha1.New().Sum(mes)
}
