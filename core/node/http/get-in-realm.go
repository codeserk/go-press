package http

import (
	"errors"
	"net/http"
	"press/common"
	"press/common/util"
	"press/core/node"

	"github.com/gorilla/mux"
)

// @Tags Node
// Gets all the nodes in the given realm
// @Summary Gets all the nodes in the given realm
// @Description Gets all the nodes in the given realm
// @ID get-nodes
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param realmId path string string "Realm ID"
// @Success 200 {array} node.Entity
// @Failure 400 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /v1/realm/{realmId}/node [get]
func getInRealm(s node.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentUser := common.GetUser(r.Context())
		if currentUser == nil {
			util.UnauthorizedError(w)
			return
		}

		params := mux.Vars(r)
		realmID := params["realmId"]
		if realmID == "" {
			util.ValidationError(w, errors.New("invalid request, realm is missing"))
			return
		}

		result, err := s.GetInRealm(realmID)
		if err != nil {
			util.InternalError(w, err)
			return
		}

		util.SendJSON(w, result)
	})
}
