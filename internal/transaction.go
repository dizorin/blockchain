package internal

type Transaction struct {
	Id         string
	Operations []Operation
	Nonce      int
}

func CreateTransaction(operations []Operation, nonce int) *Transaction {
	return &Transaction{"", operations, nonce}
}

func (t Transaction) String() string {
	return ""
}
