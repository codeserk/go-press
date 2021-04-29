package http

import (
	"press/core/field"

	"github.com/gorilla/mux"
)

func MakeHandlers(router *mux.Router, service field.Service) {
	// Create
	router.Handle("/v1/realm/{realmID}/schema/{schemaID}/field", create(service)).Methods("POST", "OPTIONS")
	router.Handle("/v1/realm/{realmID}/schema/{schemaID}/field", getBySchema(service)).Methods("GET", "OPTIONS")
}
