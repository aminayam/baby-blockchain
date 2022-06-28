package KeyPairSign

import (
	"crypto/rand"
	"crypto/rsa"
)

type KeyPair struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func GenKeyPair() KeyPair { //generates random key pair
	privateKeyStruct, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	privateKey := privateKeyStruct
	publicKey := privateKey.PublicKey
	return KeyPair{PrivateKey: privateKey, PublicKey: &publicKey}
}
