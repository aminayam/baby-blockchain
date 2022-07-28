package blockchain

type Operation struct {
	Sender    Account
	WalletID  int
	Receiver  Account
	Token     string
	Signature []byte
}

func CreateOperation(bc Blockchain, sender int, walletID int, receiver int, token string) Operation {
	senderAcc, err1 := GetAccountById(bc, sender)
	receiverAcc, err2 := GetAccountById(bc, receiver)

	if !err1 || !err2 { //stop in case wrong account id (that don`t exist)
		return Operation{}
	}

	operation := senderAcc.CreatePaymentOp(receiverAcc, walletID, token)
	return *operation
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
