package http

import (
	"net/http"
	"press/common"
	"press/common/util"
	"press/core/realm/service"
)

// @Tags Realm
// Gets all the realms accessible for the current user
// @Summary Gets all the realms accessible for the current user
// @Description Gets all the realms accessible for the current user
// @ID get-realms
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200
// @Failure 400 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /v1/realm [get]
func findByAuthor(s service.Interface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentUser := common.GetUser(r.Context())
		if currentUser == nil {
			util.UnauthorizedError(w)
			return
		}

		realms, err := s.GetForUser(currentUser)
		if err != nil {
			util.InternalError(w, err)
		}

		util.SendJSON(w, realms)
	})
}
