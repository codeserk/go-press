package http

import (
	"press/core/realm/service"

	"github.com/gorilla/mux"
)

func MakeHandlers(router *mux.Router, realmService service.Interface) {
	// Create
	router.Handle("/v1/realm", create(realmService)).Methods("POST", "OPTIONS")
	router.Handle("/v1/realm", findByAuthor(realmService)).Methods("GET", "OPTIONS")
}
