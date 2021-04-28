package service

import (
	"fmt"
	"press/core/realm"
	"press/core/realm/repository"
)

func (s *service) Create(params CreateParams) (*realm.Entity, error) {
	result, err := s.repository.CreateOne(repository.CreateOneParams{Name: params.Name, Author: params.Author})
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
