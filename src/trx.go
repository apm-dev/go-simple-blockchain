package src

type trx struct {
	Sender    string
	Recipient string
	Amount    int64
}

//	create new transaction
func NewTrx(sender string, recipient string, amount int64) *trx {
	return &trx{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
}
