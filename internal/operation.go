package internal

type Operation struct {
	Sender    Account
	Receiver  Account
	Amount    int
	Signature []byte
}

func CreateOperation(sender Account, receiver Account, amount int, signature []byte) *Operation {
	return &Operation{
		Sender:    sender,
		Receiver:  receiver,
		Amount:    amount,
		Signature: signature,
	}
}

func (o Operation) String() string {
	return ""
}
