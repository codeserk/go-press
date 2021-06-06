package http

import (
	"errors"
	"net/http"
	"press/common"
	"press/common/util"
	"press/core/field"

	"github.com/gorilla/mux"
)

// @Tags Field
// Deletes a field
// @Summary Deletes a field
// @Description Deletes a field
// @ID delete-field
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param realmId path string string "Realm ID"
// @Param schemaId path string string "Schema ID"
// @Param fieldId path string string "Field ID"
// @Success 200 {boolean} true
// @Failure 400 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /v1/realm/{realmId}/schema/{schemaId}/field/{fieldId} [delete]
func delete(s field.Service) http.Handler {
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
		if schemaID == "" {
			util.ValidationError(w, errors.New("invalid request, field is missing"))
			return
		}

		err := s.Delete(fieldID)
		if err != nil {
			util.InternalError(w, err)
			return
		}

		util.SendJSON(w, true)
	})
}
