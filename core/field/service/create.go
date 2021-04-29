package service

import (
	"fmt"
	"press/core/field"
)

func (s *service) Create(params field.CreateParams) (*field.Entity, error) {
	result, err := s.repository.InsertOne(field.InsertOneParams(params))
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
