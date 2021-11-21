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
	// ID of the schema.
	SchemaID string `json:"schemaId" validate:"required" example:"507f191e810c19729de860ea"`

	// Type of node.
	Type node.Type `json:"type" enums:"scene,model,view" validate:"oneof=scene model view"`

	// Slug of the node, used to create URIs
	Slug string `json:"slug" validate:"required" example:"how-to-write-better-go-code"`

	// Name of the node.
	Name string `json:"name" validate:"required" example:"How to write better go code"`

	// Data for the node. The structure of the data depends on the schema, and
	// it will be validated.
	Data *map[string]interface{} `json:"data"`
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
			Type:     input.Type,
			SchemaID: input.SchemaID,
			Slug:     input.Slug,
			Name:     input.Name,
			Data:     input.Data,
		})
		if err != nil {
			util.InternalError(w, err)
			return
		}

		util.SendJSON(w, result)
	})
}
