package KeyPairSign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func signData(privateKey *rsa.PrivateKey, data string) ([]byte, []byte) {
	msg := []byte(data)

	msgHash := sha256.New()
	_, err := msgHash.Write(msg)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)

	signature, err1 := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err1 != nil {
		panic(err1)
	}
	return signature, msgHashSum
}

func verifySignature(signature []byte, publicKey *rsa.PublicKey, msgHash []byte) bool {
	err := rsa.VerifyPSS(publicKey, crypto.SHA256, msgHash, signature, nil)
	if err != nil {
		fmt.Println("could not verify signature: ", err)
		return false
	}
	fmt.Println("signature verified")

	return true
}
