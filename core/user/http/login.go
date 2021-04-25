package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"press/core/user/service"
	"press/core/util"
)

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required`
} // @name LoginRequest

// @Summary Tries to login using some credentials.
// @Description Tries to login using some credentials.
// @ID login
// @Accept  json
// @Produce  json
// @Param body body loginRequest loginRequest "User login with email and password"
// @Success 200
// @Router /v1/auth/login [post]
func login(userService service.Interface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input loginRequest
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			util.ValidationError(w, errors.New("Invalid input"))
			return
		}

		err = validate.Struct(input)
		if err != nil {
			util.ValidationError(w, err)
			return
		}

		loggedUser, err := userService.Login(service.LoginParams(input))
		if err != nil {
			util.InternalError(w, err)
			return
		}

		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(loggedUser)
	})
}
