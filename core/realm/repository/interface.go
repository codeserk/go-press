package repository

import (
	"press/core/realm"
	"press/core/user"
)

type CreateOneParams struct {
	Name   string
	Author *user.Entity
}

type Interface interface {
	CreateOne(params CreateOneParams) (*realm.Entity, error)

	FindOneByID(id string) (*realm.Entity, error)
	FindByAuthor(author *user.Entity) ([]*realm.Entity, error)
}
