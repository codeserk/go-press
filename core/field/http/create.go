package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"press/common"
	"press/common/util"
	"press/core/field"
	"press/core/primitive"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type createFieldRequest struct {
	Name      string         `json:"name" validate:"required"`
	Primitive primitive.Type `json:"primitive" validate:"required"`
	Data      interface{}    `json:"data"`
} // @name CreateFieldRequest

// @Tags Field
// Creates a new field
// @Summary Creates a new field
// @Description Creates a new field
// @ID create-field
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param realmId path string string "Realm ID"
// @Param schemaId path string string "Schema ID"
// @Param body body createFieldRequest createFieldRequest "Field parameters"
// @Success 200 {object} field.Entity
// @Failure 400 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /v1/realm/{realmId}/schema/{schemaId}/field [post]
func create(s field.Service) http.Handler {
	validate := validator.New()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentUser := common.GetUser(r.Context())
		if currentUser == nil {
			util.UnauthorizedError(w)
			return
		}

		params := mux.Vars(r)
		realmID := params["realmId"]
		schemaID := params["schemaId"]
		if realmID == "" {
			util.ValidationError(w, errors.New("invalid request, realm is missing"))
			return
		}
		if schemaID == "" {
			util.ValidationError(w, errors.New("invalid request, schema is missing"))
			return
		}

		var input createFieldRequest
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

		result, err := s.Create(field.CreateParams{
			SchemaID:  schemaID,
			Name:      input.Name,
			Primitive: input.Primitive,
			Data:      input.Data,
		})
		if err != nil {
			util.InternalError(w, err)
			return
		}

		util.SendJSON(w, result)
	})
}
