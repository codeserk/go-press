package service

import (
	"net/http"
	"press/common/jwt"
	"press/core/user"
	"press/core/user/repository"
)

type RegisterParams struct {
	Email    string
	Password string
}

type LoginParams struct {
	Email    string
	Password string
}

type Interface interface {
	// Register Registers a new user with email and password. Returns the newly
	// created user and the JWT
	Register(params RegisterParams) (*user.Entity, string, error)

	// Login Logins using email and password. Returns the logged in user and the jwt.
	Login(params LoginParams) (*user.Entity, string, error)

	CreateAuthMiddleware(next http.Handler) (http.Handler)

	GenerateJWTForUser(user *user.Entity) (string ,error)
}


type service struct {
	repository repository.Interface
	jwt jwt.Interface
}

func New(repository repository.Interface, jwt jwt.Interface) Interface {
	return &service{repository, jwt}
}
