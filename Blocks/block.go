package Blocks

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type Block struct {
	BlockID       string      //уникальный идентификатор блока (хэш-значение от всех остальных данных).
	PrevBlockHash string      //идентификатор предыдущего блока (необходим для обеспечения проверки целостности истории).
	Transaction   Transaction //список транзакций, подтверждаемых в данном блоке.
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
