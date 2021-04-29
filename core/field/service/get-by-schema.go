package service

import (
	"fmt"
	"press/core/field"
)

func (s *service) GetBySchema(schemaID string) ([]*field.Entity, error) {
	result, err := s.repository.FindBySchema(schemaID)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
