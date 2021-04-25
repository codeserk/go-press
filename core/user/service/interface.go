package service

import "press/core/user"

type RegisterParams struct {
	Email    string
	Password string
}

type LoginParams struct {
	Email    string
	Password string
}

type Interface interface {
	Register(params RegisterParams) (*user.Entity, error)
	Login(params LoginParams) (*user.Entity, error)
}
