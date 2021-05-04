package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"press/common"
	"press/common/util"
	"press/core/realm/service"

	"github.com/go-playground/validator"
)

type createRequest struct {
	Name string `json:"name" validate:"required"`
}

// @Tags Realm
// Creates a new realm
// @Summary Creates a new realm
// @Description Creates a new realm
// @ID create-realm
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param body body createRequest createRequest "Realm parameters"
// @Success 200 {object} realm.Entity
// @Failure 400 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /v1/realm [post]
func create(s service.Interface) http.Handler {
	validate := validator.New()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input createRequest
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

		realms, err := s.Create(service.CreateParams{Name: input.Name, Author: currentUser})
		if err != nil {
			util.InternalError(w, err)
		}

		util.SendJSON(w, realms)
	})
}
