package service

import (
	"fmt"
	"press/core/node"
)

func (s *service) GetInRealm(realmID string) ([]*node.Entity, error) {
	result, err := s.repository.FindInRealm(realmID)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
