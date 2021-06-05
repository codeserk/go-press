package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"press/common"
	"press/common/util"
	"press/core/node"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type updateNodeRequest struct {
	Slug *string      `json:"slug"`
	Name *string      `json:"name"`
	Data *interface{} `json:"data"`
} // @name UpdateNodeRequest

// @Tags Node
// Updates a node
// @Summary Updates a node
// @Description Updates a node
// @ID update-node
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param realmId path string string "Realm ID"
// @Param nodeId path string string "Node ID"
// @Param body body updateNodeRequest updateNodeRequest "Node parameters"
// @Success 200 {object} node.Entity
// @Failure 400 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /v1/realm/{realmId}/node/{nodeId} [patch]
func update(s node.Service) http.Handler {
	validate := validator.New()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentUser := common.GetUser(r.Context())
		if currentUser == nil {
			util.UnauthorizedError(w)
			return
		}

		params := mux.Vars(r)
		realmID := params["realmId"]
		nodeID := params["nodeId"]
		if realmID == "" {
			util.ValidationError(w, errors.New("invalid request, realm is missing"))
			return
		}
		if nodeID == "" {
			util.ValidationError(w, errors.New("invalid request, node is missing"))
			return
		}

		var input updateNodeRequest
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			util.ValidationError(w, fmt.Errorf("invalid input: %v", err))
			return
		}
		err = validate.Struct(input)
		if err != nil {
			util.ValidationError(w, err)
			return
		}

		result, err := s.Update(nodeID, node.UpdateParams(input))
		if err != nil {
			util.InternalError(w, err)
			return
		}

		util.SendJSON(w, result)
	})
}
