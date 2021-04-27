package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"press/common"
	"press/common/util"
	"press/core/user"
	"press/core/user/service"
)

type meResponse struct {
	User *user.Entity `json:"user"`
	Jwt  string       `json:"jwt"`
} // @name MeResponse

// @Tags User
// Gets the current user and JWT
// @Summary Gets the current user and JWT
// @Description Gets the current user and JWT
// @ID me
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} meResponse "User and JWT"
// @Failure 400 {object} util.HttpError
// @Failure 500 {object} util.HttpError
// @Router /v1/user/me [get]
func me(userService service.Interface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := common.GetUser(r.Context())
		if user == nil {
			util.UnauthorizedError(w)
			return
		}

		jwt, err := userService.GenerateJWTForUser(user)
		if err != nil {
			util.InternalError(w, fmt.Errorf("error found while trying to genrate token for the user: %v", err))
		}
		response := loginResponse{
			User: user,
			Jwt:  jwt,
		}

		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
}
