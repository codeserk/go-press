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

type createNodeRequest struct {
	SchemaID string `json:"schemaId" validate:"required"`
	Slug     string `json:"slug" validate:"required"`
	Name     string `json:"name" validate:"required"`
} // @name CreateNodeRequest

// @Tags Node
// Creates a new node
// @Summary Creates a new node
// @Description Creates a new node
// @ID create-node
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param realmId path string string "Realm ID"
// @Param body body createNodeRequest createNodeRequest "Node parameters"
// @Success 200 {object} node.Entity
// @Failure 400 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /v1/realm/{realmId}/node [post]
func create(s node.Service) http.Handler {
	validate := validator.New()

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

		var input createNodeRequest
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

		result, err := s.Create(node.CreateParams{
			RealmID:  realmID,
			SchemaID: input.SchemaID,
			Slug:     input.Slug,
			Name:     input.Name,
		})
		if err != nil {
			util.InternalError(w, err)
			return
		}

		util.SendJSON(w, result)
	})
}
