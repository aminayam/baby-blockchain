package blockchain

type Blockchain struct {
	CoinDatabase map[int][]string //таблица отражающая текущее состояние балансов в системе. В качестве ключа используется идентификатор аккаунтa, в качестве значения - список владений пользователя.
	Blocks       []*Block         //массив хранящий все блоки добавленные в историю.
	TxDatabase   []Transaction    //массив хранящий все транзакции в истории.
	Accounts     []Account
}

func NewGenesisBlock() *Block {
	var nullTx Transaction
	resBlock := CreateBlock(nullTx, "Genesis Block")
	return resBlock
}

func InitBlockchain() *Blockchain {
	db := make(map[int][]string)
	var tx []Transaction
	genBlock := NewGenesisBlock()
	return &Blockchain{CoinDatabase: db, Blocks: []*Block{genBlock}, TxDatabase: tx}
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
	blockchain.Blocks = append(blockchain.Blocks, newBlock)

	blockchain.Accounts[sender-1].UpdateBalanceSender(token)
	blockchain.CoinDatabase[sender] = Remove(blockchain.CoinDatabase[sender], token)
	blockchain.Accounts[receiver-1].UpdateBalance(token)
	blockchain.CoinDatabase[receiver] = append(blockchain.CoinDatabase[receiver], token)
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

func verifyAccounts(bc Blockchain, sender int, receiver int) bool {
	if containsAcc(bc.Accounts, GetAccountById(bc, sender)) {
		return true
	}
	if containsAcc(bc.Accounts, GetAccountById(bc, receiver)) {
		return true
	}
	return false
}
