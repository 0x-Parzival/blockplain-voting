package router

import (
	"github.com/gorilla/mux"
	"github.com/0x-Parzival/blockplain-voting/backend/handler"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Voting endpoint
	r.HandleFunc("/vote", handler.VoteHandler).Methods("POST")

	// Results endpoint
	r.HandleFunc("/results", handler.ResultsHandler).Methods("GET")

	return r
}
