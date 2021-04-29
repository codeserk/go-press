package http

import (
	"press/core/schema"

	"github.com/gorilla/mux"
)

func MakeHandlers(router *mux.Router, service schema.Service) {
	// Create
	router.Handle("/v1/realm/{realmId}/schema", create(service)).Methods("POST", "OPTIONS")
}
