package service

import "press/core/node"

type service struct {
	repository node.Repository
}

func New(repository node.Repository) node.Service {
	return &service{repository}
}
