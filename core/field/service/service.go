package service

import (
	"press/core/field"
)

type service struct {
	repository field.Repository
}

func New(repository field.Repository) field.Service {
	return &service{repository}
}
