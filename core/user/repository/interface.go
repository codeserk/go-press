package repository

import "press/core/user"

type CreateOneParams struct {
	Email    string
	Password string
}

type Interface interface {
	CreateOne(params CreateOneParams) (*user.Entity, error)

	FindOneByID(id string) (*user.Entity, error)
	FindOneByEmail(email string) (*user.Entity, error)
}
