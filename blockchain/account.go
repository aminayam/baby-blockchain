package blockchain

import (
	"baby-blockchain/signature"
	"fmt"
)

var ID []int

type Account struct {
	AccountID int                       //unique id of account(start from '1' and increase one by one)
	Wallet    map[int]signature.KeyPair //wallet with key pairs
	Balance   []string                  //balance of account (list of current tokens that belong to account)
}

func GenAccount(bc *Blockchain) {
	wallet := make(map[int]signature.KeyPair)
	wallet[1] = signature.GenKeyPair()
	balance := make([]string, 100)
	bc.Accounts = append(bc.Accounts, Account{AccountID: idGen(), Wallet: wallet, Balance: balance})
}

func (account *Account) AddKeyPairToWallet(NewKeyPair signature.KeyPair) {
	account.Wallet[len(account.Wallet)+1] = NewKeyPair
}

func (account *Account) UpdateBalance(newItem string) {
	account.Balance = append(account.Balance, newItem)
}

func (account *Account) UpdateBalanceSender(newItem string) {
	account.Balance = Remove(account.Balance, newItem)
}

func (account *Account) CreatePaymentOp(receiver Account, id int, token string) *Operation {
	sign := account.SignData([]byte(token), id)
	return &Operation{Sender: *account, Receiver: receiver, Token: token, Signature: sign}
}

func (account *Account) GetBalance() []string {
	return account.Balance
}

func (account *Account) PrintBalance() {
	fmt.Println(account.Balance)
}

func (account *Account) SignData(data []byte, id int) []byte {
	return signature.SignData(account.Wallet[id].PrivateKey, data)
}

func idGen() int {
	newId := len(ID) + 1
	ID = append(ID, len(ID)+1)
	return newId
}

func GetAccountById(bc Blockchain, id int) (Account, bool) {
	for _, acc := range bc.Accounts {
		if acc.AccountID == id {
			return acc, true
		}
	}
	return Account{}, false
}

func Remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
