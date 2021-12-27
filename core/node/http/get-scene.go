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
// Gets a Scene by its slug
// @Summary Gets a Scene by its slug
// @Description Gets a Scene by its slug
// @ID get-scene
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param realmId path string string "Realm ID"
// @Param slug query string false "Scene Slug"
// @Success 200 {object} node.Entity
// @Failure 400 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /v1/realm/{realmId}/scene [get]
func getScene(s node.Service) http.Handler {
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
		slug := r.FormValue("slug")
		if slug == "" {
			util.ValidationError(w, errors.New("invalid request, slug is missing"))
			return
		}

		result, err := s.GetBySlug(realmID, slug)
		if err != nil {
			util.InternalError(w, err)
			return
		}

		util.SendJSON(w, result)
	})
}
