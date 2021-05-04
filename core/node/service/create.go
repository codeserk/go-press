package service

import (
	"fmt"
	"press/core/node"
)

func (s *service) Create(params node.CreateParams) (*node.Entity, error) {
	schema, err := s.schemas.GetByID(params.SchemaID)
	if err != nil {
		return nil, fmt.Errorf("invalid schema: %v", err)
	}

	insertParams := node.InsertOneParams{
		RealmID:  params.RealmID,
		SchemaID: params.SchemaID,
		Type:     params.Type,
		Slug:     params.Slug,
		Name:     params.Name,
		Data:     schema.DefaultValue(),
	}

	result, err := s.repository.InsertOne(insertParams)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
