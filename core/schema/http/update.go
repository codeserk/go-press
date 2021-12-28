package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"press/common"
	"press/common/util"
	"press/core/node"
	"press/core/schema"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type updateSchemaRequest struct {
	Type node.Type `json:"type" enums:"scene,nested"`
	Name string    `json:"name"`
} // @name UpdateSchemaRequest

// @Tags Schema
// Updates a schema
// @Summary Updates a schema
// @Description Updates a schema
// @ID update-schema
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param realmId path string string "Realm ID"
// @Param schemaId path string string "Schema ID"
// @Param body body updateSchemaRequest updateSchemaRequest "Schema parameters"
// @Success 200 {object} schema.Entity
// @Failure 400 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /v1/realm/{realmId}/schema/{schemaId} [patch]
func update(s schema.Service) http.Handler {
	validate := validator.New()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		realmID := params["realmId"]
		if realmID == "" {
			util.ValidationError(w, errors.New("invalid request, realm is missing"))
		}
		schemaID := params["schemaId"]
		if schemaID == "" {
			util.ValidationError(w, errors.New("invalid request, schema is missing"))
		}

		var input updateSchemaRequest
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			util.ValidationError(w, errors.New("invalid input"))
			return
		}
		currentUser := common.GetUser(r.Context())
		if currentUser == nil {
			util.UnauthorizedError(w)
			return
		}
		err = validate.Struct(input)
		if err != nil {
			util.ValidationError(w, err)
			return
		}

		result, err := s.Update(schemaID, schema.UpdateParams{
			Type: input.Type,
			Name: input.Name,
		})
		if err != nil {
			util.InternalError(w, err)
			return
		}

		util.SendJSON(w, result)
	})
}
