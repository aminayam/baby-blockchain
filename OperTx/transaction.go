package OperTx

type Transaction struct {
	TransactionID   string      //уникальный идентификатор транзакции (хэш-значение от всех остальных полей транзакции).
	SetOfOperations []Operation //набор операций платежей, подтверждаемых в данной транзакции.
	Nonce           int         //значение для защиты дублирования транзакций с одинаковыми операциями.
}

func createTransaction(setOfOperations []Operation, nonce int) Transaction {
	var id string

	return Transaction{TransactionID: id, SetOfOperations: setOfOperations, Nonce: nonce}
}
