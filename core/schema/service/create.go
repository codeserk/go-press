package service

import (
	"fmt"
	"press/core/schema"
)

func (s *service) Create(params schema.CreateParams) (*schema.Entity, error) {
	result, err := s.repository.InsertOne(schema.InsertOneParams(params))
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
