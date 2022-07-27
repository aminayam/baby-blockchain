package Blocks

import (
	"encoding/json"
)

var NonceList []int

type Transaction struct {
	TransactionID string    //уникальный идентификатор транзакции (хэш-значение от всех остальных полей транзакции).
	Operation     Operation //набор операций платежей, подтверждаемых в данной транзакции.
	Nonce         int       //уникальное значение, счетчик транзакций
}

func CreateTransaction(operation Operation) Transaction {
	id, _ := json.Marshal(operation)
	return Transaction{TransactionID: string(id), Operation: operation, Nonce: Nonce()}
}

func VerifyTransaction(bc Blockchain, tx Transaction) bool {
	if containsTx(bc.TxDatabase, tx) { //проверка, не была ли уже ли транзакция в истории
		return true
	}
	return false
}

func containsTx(s []Transaction, e Transaction) bool {
	for _, a := range s {
		if a.TransactionID == e.TransactionID && a.Nonce == e.Nonce {
			return false
		}
	}
	return true
}

func Nonce() int {
	newNonce := len(NonceList) + 1
	NonceList = append(NonceList, newNonce)
	return newNonce
}
