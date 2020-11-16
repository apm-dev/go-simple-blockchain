package bc

//	Caution: Do not change fields name and order, because of hash algorithm logic
type block struct {
	Index        int32
	Timestamp    int64
	Trxs         []*trx
	Nonce        int64
	PreviousHash string
}
