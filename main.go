package main

import (
	"baby-blockchain/blockchain"
	"baby-blockchain/signature"
	"fmt"
)

func main() {
	bc := blockchain.InitBlockchain() //we receive firs genesis block & null account

	///migrations -> generates 3 accounts with tokens
	blockchain.GenAccount(bc)                                 //generates account
	bc.Accounts[1].AddKeyPairToWallet(signature.GenKeyPair()) //generates new key-pair for acc1
	blockchain.GenAccount(bc)
	blockchain.GenAccount(bc)

	bc.CoinDatabase[1] = []string{"token1", "token2", "token3", "diamond"}
	bc.CoinDatabase[2] = []string{"the_best_token_ever", "abcd"}
	bc.CoinDatabase[3] = []string{"cool_thing"}

	bc.Accounts[1].Balance = []string{"token1", "token2", "token3"}
	bc.Accounts[2].Balance = []string{"the_best_token_ever", "abcd"}
	bc.Accounts[3].Balance = []string{"cool_thing"}
	///end of migrations
	printAccounts(bc)

	//let`s start to use our blockchain!

	//we can add block with tx - AddBlock(senderId int, walletID(key^ that you choose to sign data) int, receiverID int, token_that_you_want_send string)
	bc.AddBlock(1, 2, 2, "token1")
	bc.AddBlock(2, 1, 3, "the_best_token_ever")
	bc.AddBlock(3, 1, 2, "cool_thing")

	bc.AddBlock(1, 1, 2, "lalala") //if we try to add wrong data(token that don't exist), block will not be added to blockchain
	bc.AddBlock(5, 1, 8, "token2") //if we try to add wrong data(account that don't exist), block will not be added to blockchain

	//now we can check blockchain and it`s databases to see what happened
	printBlockchain(bc)

	//Also we can see current balances of accounts
	printAccounts(bc)
	//or for each account
	fmt.Println("Current balance of 3rd account:")
	bc.Accounts[3].PrintBalance()

	//let`s check signification mechanism VerifySignature
	//VerifyTransactionSignature(bc Blockchain,txNonce int) Nonce = index number of tx
	ans := blockchain.VerifyTransactionSignature(bc, 3)

	if ans {
		fmt.Println("signature verified!")
	} else {
		fmt.Println("sorry, but no(")
	}
}

func printBlocks(bc *blockchain.Blockchain) {
	for _, block := range bc.Blocks {
		fmt.Println("")
		fmt.Printf("Previous Hash: %x\n", block.PrevBlockHash)
		fmt.Println("Data in Block:", //block.Transaction.TransactionID,
			"/Nonce - ", block.Transaction.Nonce,
			"/SenderId - ", block.Transaction.Operation.Sender.AccountID,
			"/ReceiverID - ", block.Transaction.Operation.Receiver.AccountID,
			//"/Signature", block.Transaction.Operation.Signature,
			"/Token - ", block.Transaction.Operation.Token)
		fmt.Printf("Hash: %x\n", block.BlockID)
	}
}

func printAccounts(bc *blockchain.Blockchain) {
	for _, acc := range bc.Accounts {
		fmt.Println(acc)
	}
}

func printBlockchain(bc *blockchain.Blockchain) {
	fmt.Println("Tokens database:", bc.CoinDatabase)
	//fmt.Println("Transactions database:", bc.TxDatabase)
	fmt.Println("Accounts:")
	printAccounts(bc)
	fmt.Println("Blocks:")
	printBlocks(bc)

}
