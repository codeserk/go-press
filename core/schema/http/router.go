package http

import (
	"press/core/schema"

	"github.com/gorilla/mux"
)

func MakeHandlers(router *mux.Router, service schema.Service) {
	router.Handle("/v1/realm/{realmId}/schema", create(service)).Methods("POST", "OPTIONS")
	router.Handle("/v1/realm/{realmId}/schema/{schemaId}", update(service)).Methods("PATCH", "OPTIONS")
	router.Handle("/v1/realm/{realmId}/schema", getInRealm(service)).Methods("GET", "OPTIONS")
	router.Handle("/v1/realm/{realmId}/schema/{schemaId}", delete(service)).Methods("DELETE", "OPTIONS")
}
