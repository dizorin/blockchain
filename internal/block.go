package internal

type Block struct {
	Id           int
	PrevHash     string
	Transactions []Transaction
}

func CreateBlock(transactions []Transaction, prevHash string) *Block {
	return &Block{
		Id:           0,
		PrevHash:     prevHash,
		Transactions: transactions,
	}
}

func (b Block) String() string {
	return ""
}
