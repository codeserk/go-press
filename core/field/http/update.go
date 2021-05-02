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

type updateFieldRequest struct {
	Key       *string         `json:"key"`
	Name      *string         `json:"name"`
	Primitive *primitive.Type `json:"primitive"`
	Config    *interface{}    `json:"config"`
} // @name UpdateFieldRequest

// @Tags Field
// Updates a field
// @Summary Updates a field
// @Description Updates a field
// @ID update-field
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param realmId path string string "Realm ID"
// @Param schemaId path string string "Schema ID"
// @Param fieldId path string string "Field ID"
// @Param body body updateFieldRequest updateFieldRequest "Field parameters"
// @Success 200 {object} field.Entity
// @Failure 400 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /v1/realm/{realmId}/schema/{schemaId}/field/{fieldId} [patch]
func update(s field.Service) http.Handler {
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
		fieldID := params["fieldId"]
		if realmID == "" {
			util.ValidationError(w, errors.New("invalid request, realm is missing"))
			return
		}
		if schemaID == "" {
			util.ValidationError(w, errors.New("invalid request, schema is missing"))
			return
		}
		if fieldID == "" {
			util.ValidationError(w, errors.New("invalid request, field is missing"))
			return
		}

		var input updateFieldRequest
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

		result, err := s.Update(fieldID, field.UpdateParams(input))
		if err != nil {
			util.InternalError(w, err)
			return
		}

		util.SendJSON(w, result)
	})
}
