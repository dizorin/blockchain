package internal

type Blockchain struct {
	CoinDatabase map[int]int
	BlockHistory []Block
	TxDatabase   []Transaction
	FaucetCoins  int
}

func InitBlockchain() *Blockchain {
	return nil
}

func (b *Blockchain) getTokenFromFaucet(account Account, amount int) {
}

func (b *Blockchain) validateBlock(block Block) {

}

func (b *Blockchain) ShowCoinDatabase() {

}

func (b *Blockchain) String() string {
	return ""
}
