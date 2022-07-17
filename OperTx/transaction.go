package OperTx

type Transaction struct {
	TransactionID string    //уникальный идентификатор транзакции (хэш-значение от всех остальных полей транзакции).
	Operation     Operation // операция платежей, подтверждаемых в данной транзакции.
	Nonce         int       //значение для защиты дублирования транзакций с одинаковыми операциями.
}

func CreateTransaction(Operation Operation, nonce int) Transaction {
	var id string

	return Transaction{TransactionID: id, Operation: Operation, Nonce: nonce}
}
