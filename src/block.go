package src

import "time"

type block struct {
	Index        int32
	Timestamp    time.Time
	Trxs         []*trx
	Proof        string
	PreviousHash string
}
