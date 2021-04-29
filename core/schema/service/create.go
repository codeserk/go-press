package service

import (
	"fmt"
	"press/core/schema"
)

func (s *service) Create(params schema.CreateParams) (*schema.Entity, error) {
	var fields []*schema.InsertOneFieldParams
	for _, f := range params.Fields {
		fields = append(fields, &schema.InsertOneFieldParams{Name: f.Name, Primitive: f.Primitive, Data: f.Data})
	}
	result, err := s.repository.InsertOne(schema.InsertOneParams{
		RealmID:  params.RealmID,
		AuthorID: params.AuthorID,
		Name:     params.Name,
		Fields:   fields,
	})
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
