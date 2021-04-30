package service

import (
	"fmt"
	"press/core/schema"
)

func (s *service) GetInRealm(realmID string) ([]*schema.Entity, error) {
	result, err := s.repository.FindInRealm(realmID)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
