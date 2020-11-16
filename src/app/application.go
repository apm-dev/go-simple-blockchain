package app

import (
	"github.com/apm-dev/go-simple-blockchain/src/bc"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var BC = bc.NewBlockchain()

func StartApplication() {
	r := mux.NewRouter()
	registerRoutes(r)

	log.Fatal(http.ListenAndServe(":8000", r))
}
