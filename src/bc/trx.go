package bc

type Trx struct {
	Sender    string
	Recipient string
	Amount    int64
}

//	create new transaction
func NewTrx(sender string, recipient string, amount int64) *Trx {
	return &Trx{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
}
