package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"press/common/util"
	"press/core/user"
	"press/core/user/service"
)

type loginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"test@test.com"`
	Password string `json:"password" validate:"required" example:"test"`
} // @name LoginRequest

type loginResponse struct {
	User *user.Entity `json:"user"`
	Jwt  string       `json:"jwt"`
} // @name LoginResponse

// @Tags Auth
// @Summary Tries to login using some credentials.
// @Description Tries to login using some credentials.
// @ID login
// @Accept  json
// @Produce  json
// @Param body body loginRequest loginRequest "User login with email and password"
// @Success 200 {object} loginResponse "User response and JWT"
// @Failure 400 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /v1/auth/login [post].
func login(userService service.Interface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input loginRequest
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			util.ValidationError(w, errors.New("invalid input"))
			return
		}

		err = validate.Struct(input)
		if err != nil {
			util.ValidationError(w, err)
			return
		}

		loggedUser, jwt, err := userService.Login(service.LoginParams(input))
		if err != nil {
			util.InternalError(w, err)
			return
		}
		response := loginResponse{
			User: loggedUser,
			Jwt:  jwt,
		}

		util.SendJSON(w, response)
	})
}
