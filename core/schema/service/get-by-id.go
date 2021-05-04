package service

import (
	"fmt"
	"press/core/schema"
)

func (s *service) GetByID(schemaID string) (*schema.Entity, error) {
	result, err := s.repository.FindOneByID(schemaID)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
