package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"press/common/util"
	"press/core/user"
	"press/core/user/service"

	"github.com/go-playground/validator"
)

type registerRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
} // @name RegisterRequest

type registerResponse struct {
	User *user.Entity `json:"user"`
	Jwt  string       `json:"jwt"`
} // @name RegisterResponse

var validate *validator.Validate = validator.New()

// Registers a new user with email and password
// @Summary Registers a new user with email and password
// @Description Registers a new user with email and password
// @ID register-user
// @Accept  json
// @Produce  json
// @Param body body registerRequest registerRequest "User registration parameters"
// @Success 200
// @Router /v1/auth/register [post]
func register(userService service.Interface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input registerRequest
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

		createdUser, jwt, err := userService.Register(service.RegisterParams(input))
		if err != nil {
			util.InternalError(w, err)
			return
		}
		response := registerResponse{
			User: createdUser,
			Jwt:  jwt,
		}

		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
}
