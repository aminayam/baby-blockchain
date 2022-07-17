package Blocks

import (
	"baby-blockchain/OperTx"
)

type Blockchain struct {
	Blocks []*Block //массив хранящий все блоки добавленные в историю.
}

func NewGenesisBlock() *Block {
	var nullTx OperTx.Transaction
	resBlock := CreateBlock(nullTx, "Genesis Block")
	return resBlock
}

func InitBlockchain() Blockchain {
	genBlock := NewGenesisBlock()
	return Blockchain{[]*Block{genBlock}}
}

func ValidateBlock(newBlock, oldBlock *Block) bool {

	if oldBlock.BlockID != newBlock.PrevBlockHash {
		return false
	}

	return true

}

func (blockchain *Blockchain) AddBlock(data OperTx.Transaction) {
	prevBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.BlockID)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}
