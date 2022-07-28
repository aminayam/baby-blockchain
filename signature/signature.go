package signature

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func SignData(privateKey *rsa.PrivateKey, data []byte) []byte { //RSA signing, returns signature and the result of hashing the input message using sha-256

	msgHash := sha256.New()
	_, err := msgHash.Write(data)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)

	signature, err1 := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err1 != nil {
		panic(err1)
	}
	return signature
}

func VerifySignature(signature []byte, publicKey *rsa.PublicKey, data []byte) bool { //RSA verification
	err := rsa.VerifyPSS(publicKey, crypto.SHA256, data, signature, nil)

	if err != nil { //couldn't verify signature
		return false
	} else { //signature verified
		return true
	}
}
