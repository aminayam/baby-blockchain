package KeyPairSign

import (
	"crypto/rand"
	"crypto/rsa"
)

type KeyPair struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func genKeyPair() KeyPair {
	privateKeyStruct, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	privateKey := privateKeyStruct
	publicKey := privateKey.PublicKey
	return KeyPair{privateKey, &publicKey}
}
