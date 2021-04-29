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
// Gets all the fields of the given schema
// @Summary Gets all the fields of the given schema
// @Description Gets all the fields of the given schema
// @ID get-fields
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param realmId path string string "Realm ID"
// @Param schemaId path string string "Schema ID"
// @Success 200 {array} field.Entity
// @Failure 400 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /v1/realm/{realmId}/schema/{schemaId}/field [get]
func getBySchema(s field.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentUser := common.GetUser(r.Context())
		if currentUser == nil {
			util.UnauthorizedError(w)
			return
		}

		params := mux.Vars(r)
		realmID := params["realmID"]
		schemaID := params["schemaID"]
		if realmID == "" {
			util.ValidationError(w, errors.New("invalid request, realm is missing"))
		}
		if schemaID == "" {
			util.ValidationError(w, errors.New("invalid request, schema is missing"))
		}

		result, err := s.GetBySchema(schemaID)
		if err != nil {
			util.InternalError(w, err)
		}

		util.SendJSON(w, result)
	})
}
