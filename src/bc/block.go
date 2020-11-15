package bc

//	Caution: Do not change fields name and order, because of hash algorithm logic
type block struct {
	Index        int32
	Timestamp    int64
	Trxs         []*Trx
	Nonce        int64
	PreviousHash string
}

func NewBlock(i int32, ts int64, trxs []*Trx, ph string) *block {
	return &block{
		Index:        i,
		Timestamp:    ts,
		Trxs:         trxs,
		Nonce:        0,
		PreviousHash: ph,
	}
}
