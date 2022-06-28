package OperTx

import "baby-blockchain/Account"

type Operation struct {
	Sender    Account.Account
	Receiver  Account.Account
	Items     []string
	Signature []byte
}

func createOperation(sender Account.Account, receiver Account.Account, items []string, signature []byte) {
}

func verifyOperation(operation Operation) bool {
	return true
}
