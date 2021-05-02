package service

import (
	"fmt"
	"press/core/node"
)

func (s *service) Create(params node.CreateParams) (*node.Entity, error) {
	result, err := s.repository.InsertOne(node.InsertOneParams(params))
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
