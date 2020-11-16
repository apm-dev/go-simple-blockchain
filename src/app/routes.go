package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func registerRoutes(r *mux.Router) {

	r.HandleFunc("/chain", GetChainHandler).Methods(http.MethodGet)
	r.HandleFunc("/mine", MineHandler).Methods(http.MethodGet)
	r.HandleFunc("/trxs/new", NewTrxHandler).Methods(http.MethodPost)

}
