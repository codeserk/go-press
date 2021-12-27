package http

import (
	"press/core/node"

	"github.com/gorilla/mux"
)

func MakeHandlers(router *mux.Router, service node.Service) {
	router.Handle("/v1/realm/{realmId}/scene", getScene(service)).Queries("slug", "{slug}").Methods("GET", "OPTIONS")
	router.Handle("/v1/realm/{realmId}/node", getInRealm(service)).Methods("GET", "OPTIONS")
	router.Handle("/v1/realm/{realmId}/node", create(service)).Methods("POST", "OPTIONS")
	router.Handle("/v1/realm/{realmId}/node/{nodeId}", update(service)).Methods("PATCH", "OPTIONS")
}
