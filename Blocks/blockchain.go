package Blocks

import (
	"baby-blockchain/Account"
	"baby-blockchain/OperTx"
)

type Blockchain struct {
	CoinDatabase map[string][]string  //таблица отражающая текущее состояние балансов в системе. В качестве ключа используется идентификатор аккаунт, в качестве значения - список владений пользователя.
	Blocks       []*Block             //массив хранящий все блоки добавленные в историю.
	TxDatabase   []OperTx.Transaction //массив хранящий все транзакции в истории.
}

func NewGenesisBlock() *Block {
	var nullTx []OperTx.Transaction
	resBlock := CreateBlock(nullTx, "Genesis Block")
	return resBlock
}

func InitBlockchain() *Blockchain {
	var db map[string][]string
	var tx []OperTx.Transaction
	genBlock := NewGenesisBlock()
	return &Blockchain{CoinDatabase: db, Blocks: []*Block{genBlock}, TxDatabase: tx}
}

func (blockchain Blockchain) GetTokenFromFaucet(account Account.Account, amount int) {

}

func ValidateBlock() {

}

func (blockchain *Blockchain) AddBlock(data []OperTx.Transaction) {
	prevBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.BlockID)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}
