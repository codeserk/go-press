package http

import (
	"fmt"
	"net/http"
	"press/common"
	"press/common/util"
	"press/core/user"
	"press/core/user/service"
)

// asdas.
type meResponse struct {
	User *user.Entity `json:"user" validate:"required"`
	Jwt  string       `json:"jwt" validate:"required"`
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
// @Failure 400 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /v1/user/me [get].
func me(userService service.Interface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentUser := common.GetUser(r.Context())
		if currentUser == nil {
			util.UnauthorizedError(w)
			return
		}

		jwt, err := userService.GenerateJWTForUser(currentUser)
		if err != nil {
			util.InternalError(w, fmt.Errorf("error found while trying to genrate token for the user: %v", err))
		}
		response := meResponse{
			User: currentUser,
			Jwt:  jwt,
		}

		util.SendJSON(w, response)
	})
}
