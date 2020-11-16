package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

//	Get entire chain
func GetChainHandler(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(BC.GetChain())
	w.Write(j)
}

//	Mine a block
func MineHandler(w http.ResponseWriter, r *http.Request) {
	err := BC.Mine("apm-wallet")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message":%s}`, err.Error())))
		return
	}
	w.Write([]byte(`{"message":"new block mined"}`))
}

type NewTrxReq struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
}

//	Create new trx
func NewTrxHandler(w http.ResponseWriter, r *http.Request) {
	//	We assume that you send correct data
	//	We focus on blockchain not the restful api
	trxReq := NewTrxReq{}
	json.NewDecoder(r.Body).Decode(&trxReq)
	amount, _ := strconv.ParseFloat(trxReq.Amount, 32)
	BC.NewTrx(trxReq.Sender, trxReq.Recipient, float32(amount))
	w.Header().Set("Content-Type","application/json")
	w.Write([]byte(`{"message":"your trx added to mempool and will get done soon"}`))
}
