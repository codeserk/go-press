package service

import "press/core/user/repository"

type service struct {
	repository repository.Interface
}

func Create(repository repository.Interface) Interface {
	return &service{repository: repository}
}
