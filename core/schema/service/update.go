package service

import (
	"fmt"
	"press/core/schema"
)

// Update updates one schema
func (s *service) Update(id string, params schema.UpdateParams) (*schema.Entity, error) {
	result, err := s.repository.UpdateOne(id, schema.UpdateOneParams(params))
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return result, nil
}
