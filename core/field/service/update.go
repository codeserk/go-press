package service

import (
	"fmt"
	"press/core/field"
)

func (s *service) Update(fieldID string, params field.UpdateParams) (*field.Entity, error) {
	result, err := s.repository.PatchOne(fieldID, field.PatchOneParams(params))
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
