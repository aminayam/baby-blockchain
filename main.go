package main

import (
	"baby-blockchain/Blocks"
	"fmt"
)

func main() {
	blockchain := Blocks.InitBlockchain()

	///migrations -> 3 accounts with tokens
	Blocks.GenAccount(blockchain)
	Blocks.GenAccount(blockchain)
	Blocks.GenAccount(blockchain)

	fmt.Println(blockchain)

	blockchain.CoinDatabase[1] = []string{"token1", "token2", "token3"}
	blockchain.CoinDatabase[2] = []string{"the_best_token_ever", "abcd"}
	blockchain.CoinDatabase[3] = []string{"cool_thing"}

	blockchain.Accounts[0].Balance = []string{"token1", "token2", "token3"}
	blockchain.Accounts[1].Balance = []string{"the_best_token_ever", "abcd"}
	blockchain.Accounts[2].Balance = []string{"cool_thing"}
	///

	blockchain.AddBlock(1, 1, 2, "token1")
	fmt.Println(blockchain)
	blockchain.AddBlock(1, 1, 2, "lalala") //if we try to add wrong data(token that don't exist),
	fmt.Println(blockchain)

	for _, block := range blockchain.Blocks {
		fmt.Println("/////////////////////////////")
		fmt.Printf("Previous Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data in Block")
		fmt.Printf("Hash: %x\n", block.BlockID)
	}
}
