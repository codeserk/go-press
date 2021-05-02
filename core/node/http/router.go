package http

import (
	"press/core/node"

	"github.com/gorilla/mux"
)

func MakeHandlers(router *mux.Router, service node.Service) {
	router.Handle("/v1/realm/{realmId}/node", create(service)).Methods("POST", "OPTIONS")
}
