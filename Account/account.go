package Account

import (
	"baby-blockchain/KeyPairSign"
	"baby-blockchain/OperTx"
	"fmt"
	"strings"
)

var ID []int

type Account struct {
	AccountID int
	Wallet    map[int]KeyPairSign.KeyPair
	Balance   []string
}

func GenAccount() Account {
	var wallet map[int]KeyPairSign.KeyPair
	wallet[1] = KeyPairSign.GenKeyPair()
	var balance []string
	return Account{AccountID: idGen(), Wallet: wallet, Balance: balance}
}

func (account Account) AddKeyPairToWallet(NewKeyPair KeyPairSign.KeyPair) {
	account.Wallet[len(account.Wallet)+1] = NewKeyPair
}

func (account Account) UpdateBalance(newItem string) {
	account.Balance = append(account.Balance, newItem)
}

func (account Account) CreatePaymentOp(receiver Account, id int, items []string) *OperTx.Operation {
	data := strings.Join(items, " ")
	signature := account.SignData(data, id)
	return &OperTx.Operation{Sender: account, Receiver: receiver, Items: items, Signature: signature}
}

func (account Account) GetBalance() []string {
	return account.Balance
}

func (account Account) PrintBalance() {
	fmt.Println(account.Balance)
}

func (account Account) SignData(data string, id int) []byte {
	signature, _ := KeyPairSign.SignData(account.Wallet[id].PrivateKey, data)
	return signature
}

func idGen() int {
	newId := len(ID) + 1
	ID = append(ID, len(ID)+1)
	return newId
}
