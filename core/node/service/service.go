package service

import (
	"press/core/node"
	"press/core/schema"
)

type service struct {
	repository node.Repository
	schemas    schema.Service
}

func New(repository node.Repository, schemas schema.Service) node.Service {
	return &service{repository, schemas}
}
