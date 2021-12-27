package service

import (
	"fmt"
	"press/core/node"
)

func (s *service) GetBySlug(realmID string, slug string) (*node.Entity, error) {
	result, err := s.repository.FindBySlug(realmID, slug)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
