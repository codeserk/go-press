package http

import (
	"errors"
	"net/http"
	"press/common"
	"press/common/util"
	"press/core/schema"

	"github.com/gorilla/mux"
)

// @Tags Schema
// Deletes a schema by its id
// @Summary Deletes a schema by its id
// @Description Deletes a schema by its id
// @ID delete-schema
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param realmId path string string "Realm ID"
// @Param schemaId path string string "Schema ID"
// @Success 200
// @Failure 400 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /v1/realm/{realmId}/schema/{schemaId} [delete]
func delete(s schema.Service) http.Handler {
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

		currentUser := common.GetUser(r.Context())
		if currentUser == nil {
			util.UnauthorizedError(w)
			return
		}

		err := s.Delete(schemaID)
		if err != nil {
			util.InternalError(w, err)
			return
		}

		util.SendJSON(w, true)
	})
}
