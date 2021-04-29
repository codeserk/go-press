package service

import "press/core/schema"

type service struct {
	repository schema.Repository
}

func New(repository schema.Repository) schema.Service {
	return &service{repository}
}
