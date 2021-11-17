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

	var data map[string]interface{}
	if params.Data == nil {
		data = schema.DefaultValue()
	} else {
		data = *params.Data
	}
	insertParams := node.InsertOneParams{
		RealmID:  params.RealmID,
		SchemaID: params.SchemaID,
		Type:     params.Type,
		Slug:     params.Slug,
		Name:     params.Name,
		Data:     data,
	}

	result, err := s.repository.InsertOne(insertParams)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
