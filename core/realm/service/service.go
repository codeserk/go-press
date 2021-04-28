package service

import (
	"press/core/realm"
	"press/core/realm/repository"
	"press/core/user"
)

type CreateParams struct {
	Name   string
	Author *user.Entity
}

type Interface interface {
	Create(params CreateParams) (*realm.Entity, error)

	GetByID(id string) (*realm.Entity, error)
	GetForUser(user *user.Entity) ([]*realm.Entity, error)
}

type service struct {
	repository repository.Interface
}

func New(repository repository.Interface) Interface {
	return &service{repository}
}
