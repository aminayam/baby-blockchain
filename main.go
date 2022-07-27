package main

import (
	"baby-blockchain/blockchain"
	"fmt"
)

func main() {
	bc := blockchain.InitBlockchain()

	///migrations -> 3 accounts with tokens
	blockchain.GenAccount(bc)
	blockchain.GenAccount(bc)
	blockchain.GenAccount(bc)

	fmt.Println(bc)

	bc.CoinDatabase[1] = []string{"token1", "token2", "token3"}
	bc.CoinDatabase[2] = []string{"the_best_token_ever", "abcd"}
	bc.CoinDatabase[3] = []string{"cool_thing"}

	bc.Accounts[0].Balance = []string{"token1", "token2", "token3"}
	bc.Accounts[1].Balance = []string{"the_best_token_ever", "abcd"}
	bc.Accounts[2].Balance = []string{"cool_thing"}
	///

	bc.AddBlock(1, 1, 2, "token1")
	fmt.Println(bc)
	bc.AddBlock(1, 1, 2, "lalala") //if we try to add wrong data(token that don't exist),
	fmt.Println(bc)

	for _, block := range bc.Blocks {
		fmt.Println("/////////////////////////////")
		fmt.Printf("Previous Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data in Block")
		fmt.Printf("Hash: %x\n", block.BlockID)
	}
}
