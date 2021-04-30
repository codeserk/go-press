package http

import (
	"press/core/field"

	"github.com/gorilla/mux"
)

func MakeHandlers(router *mux.Router, service field.Service) {
	// Create
	router.Handle("/v1/realm/{realmId}/schema/{schemaId}/field", create(service)).Methods("POST", "OPTIONS")
	router.Handle("/v1/realm/{realmId}/schema/{schemaId}/field", getBySchema(service)).Methods("GET", "OPTIONS")
}
