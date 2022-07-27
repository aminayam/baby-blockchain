package Blocks

import (
	"baby-blockchain/KeyPairSign"
	"fmt"
)

var ID []int

type Account struct {
	AccountID int
	Wallet    map[int]KeyPairSign.KeyPair
	Balance   []string
}

func GenAccount(bc *Blockchain) {
	wallet := make(map[int]KeyPairSign.KeyPair)
	wallet[1] = KeyPairSign.GenKeyPair()
	balance := make([]string, 100)
	bc.Accounts = append(bc.Accounts, Account{AccountID: idGen(), Wallet: wallet, Balance: balance})
}

func (account *Account) AddKeyPairToWallet(NewKeyPair KeyPairSign.KeyPair) {
	account.Wallet[len(account.Wallet)+1] = NewKeyPair
}

func (account *Account) UpdateBalance(newItem string) {
	account.Balance = append(account.Balance, newItem)
}

func (account *Account) UpdateBalanceSender(newItem string) {
	account.Balance = Remove(account.Balance, newItem)
}

func (account *Account) CreatePaymentOp(receiver Account, id int, token string) *Operation {
	signature := account.SignData(token, id)
	return &Operation{Sender: *account, Receiver: receiver, Token: token, Signature: signature}
}

func (account *Account) GetBalance() []string {
	return account.Balance
}

func (account *Account) PrintBalance() {
	fmt.Println(account.Balance)
}

func (account *Account) SignData(data string, id int) []byte {
	signature, _ := KeyPairSign.SignData(account.Wallet[id].PrivateKey, data)
	return signature
}

func idGen() int {
	newId := len(ID) + 1
	ID = append(ID, len(ID)+1)
	return newId
}

func GetAccountById(bc Blockchain, id int) Account {
	for _, acc := range bc.Accounts {
		if acc.AccountID == id {
			return acc
		}
	}
	return Account{}
}

func Remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
