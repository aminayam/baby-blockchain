package OperTx

type Operation struct {
	Sender    Account
	Receiver  Account
	Items     []string
	Signature []byte
}

func createOperation(sender Account, receiver Account, items []string, signature []byte) Operation {
	return Operation{Sender: sender, Receiver: receiver, Items: items, Signature: signature}

}

func verifyOperation(operation Operation) bool {
	return true
}
