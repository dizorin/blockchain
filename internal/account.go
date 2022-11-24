package internal

import "github.com/dizorin/blockchain/internal/crypto"

type Account struct {
	Id      string
	Wallet  []crypto.KeyPair
	balance int
}

func CreateAccount() *Account {
	return nil
}

func (a *Account) addKeyPairToWallet(pair crypto.KeyPair) {
	a.Wallet = append(a.Wallet, pair)
}

func (a *Account) updateBalance() {

}

func (a *Account) signData(mes string, index int) []byte {
	return nil
}

func (a *Account) createOperation(recipient Account, amount int, index int) *Operation {
	return nil
}

func (a Account) String() string {
	return ""
}
