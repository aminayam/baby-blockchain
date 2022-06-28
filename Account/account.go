package Account

import (
	"baby-blockchain/KeyPairSign"
	"baby-blockchain/OperTx"
	"fmt"
	"strings"
)

var ID []int

type Account struct {
	accountID int
	wallet    map[int]KeyPairSign.KeyPair
	balance   []string
}

func genAccount() Account {
	var wallet map[int]KeyPairSign.KeyPair
	wallet[1] = KeyPairSign.GenKeyPair()
	var balance []string
	return Account{accountID: idGen(), wallet: wallet, balance: balance}
}

func (account Account) addKeyPairToWallet(NewKeyPair KeyPairSign.KeyPair) {
	account.wallet[len(account.wallet)+1] = NewKeyPair
}

func (account Account) updateBalance(newItem string) {
	account.balance = append(account.balance, newItem)
}

func (account Account) createPaymentOp(receiver Account, id int, items []string) *OperTx.Operation {
	data := strings.Join(items, " ")
	signature := account.signData(data, id)
	return &OperTx.Operation{Sender: account, Receiver: receiver, Items: items, Signature: signature}
}

func (account Account) getBalance() []string {
	return account.balance
}

func (account Account) printBalance() {
	fmt.Println(account.balance)
}

func (account Account) signData(data string, id int) []byte {
	signature, _ := KeyPairSign.SignData(account.wallet[id].PrivateKey, data)
	return signature
}

func idGen() int {
	newId := len(ID) + 1
	ID = append(ID, len(ID)+1)
	return newId
}
