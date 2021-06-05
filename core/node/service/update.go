package service

import (
	"fmt"
	"press/core/node"
)

func (s *service) Update(id string, params node.UpdateParams) (*node.Entity, error) {
	result, err := s.repository.PatchOne(id, node.PatchOneParams(params))
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
