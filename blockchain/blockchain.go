package blockchain

type Blockchain struct {
	CoinDatabase map[int][]string //table with current state of balances in the system. The key is the account ID, and the value is the list of the user's tokens
	Blocks       []*Block         //slice with all approved blocks in history
	TxDatabase   []Transaction    //slice with all tx`s in history
	Accounts     []Account        //slice with all accounts in system
}

func NewGenesisBlock() *Block { //generates genesis block
	var nullTx Transaction
	resBlock := CreateBlock(nullTx, "Genesis Block")
	return resBlock
}

func InitBlockchain() *Blockchain { //initialises blockchain
	db := make(map[int][]string)
	var tx []Transaction
	genBlock := NewGenesisBlock()
	return &Blockchain{CoinDatabase: db, Blocks: []*Block{genBlock}, TxDatabase: tx, Accounts: []Account{{0, nil, nil}}}
}

func ValidateBlock(newBlock, oldBlock *Block) bool {
	if oldBlock.BlockID != newBlock.PrevBlockHash {
		return false
	}
	return true
}

func (blockchain *Blockchain) AddBlock(sender int, walletID int, receiver int, token string) bool {

	verifyAccounts(*blockchain, sender, receiver)

	op := CreateOperation(*blockchain, sender, walletID, receiver, token)
	err := VerifyOperation(op)
	if !err {
		return err
	}

	tx := CreateTransaction(op)
	err = VerifyTransaction(*blockchain, tx)
	if !err {
		return err
	}

	prevBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := CreateBlock(tx, prevBlock.BlockID)

	err = ValidateBlock(newBlock, prevBlock)
	if !err {
		return err
	}

	//adding block to blockchain database after all verifications & validations
	blockchain.Blocks = append(blockchain.Blocks, newBlock)

	blockchain.Accounts[sender].UpdateBalanceSender(token)
	blockchain.Accounts[receiver].UpdateBalance(token)

	blockchain.CoinDatabase[sender] = Remove(blockchain.CoinDatabase[sender], token)
	blockchain.CoinDatabase[receiver] = append(blockchain.CoinDatabase[receiver], token)

	blockchain.TxDatabase = append(blockchain.TxDatabase, tx)

	return true
}

func containsAcc(s []Account, e Account) bool {
	for _, a := range s {
		if a.AccountID == e.AccountID {
			return false
		}
	}
	return true
}

func verifyAccounts(bc Blockchain, sender int, receiver int) bool { //checks is account exist in blockchain
	senderAcc, _ := GetAccountById(bc, sender)
	if containsAcc(bc.Accounts, senderAcc) {
		return true
	}
	recAcc, _ := GetAccountById(bc, receiver)
	if containsAcc(bc.Accounts, recAcc) {
		return true
	}
	return false
}
