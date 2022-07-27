package Blocks

type Operation struct {
	Sender    Account
	WalletID  int
	Receiver  Account
	Token     string
	Signature []byte
}

func CreateOperation(bc Blockchain, sender int, walletID int, receiver int, token string) Operation {
	return Operation{GetAccountById(bc, sender), walletID, GetAccountById(bc, receiver), token, []byte{}}
}

func VerifyOperation(op Operation) bool {
	if contains(op.Sender.Balance, op.Token) {
		return true
	}
	return false
	//fixme: add verifying by signature
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
