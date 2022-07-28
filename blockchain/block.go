package blockchain

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type Block struct {
	BlockID       string      //unique block identifier (hash value from all other data)
	PrevBlockHash string      //identifier of the previous block (required to ensure the integrity of the history)
	Transaction   Transaction //transaction that confirmed in this block
}

func CreateBlock(data Transaction, prevBlockHash string) *Block {
	var block = Block{"", prevBlockHash, data}
	err := block.CalculateHash()
	if err != nil {
		panic(err)
	}
	return &block
}
func (block *Block) CalculateHash() error {
	data, err := json.Marshal(block)
	if err != nil {
		return err
	}
	hash := md5.Sum(data)
	block.BlockID = hex.EncodeToString(hash[:])
	return nil
}
