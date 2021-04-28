package service

import (
	"fmt"
	"press/core/realm"
)

func (s *service) GetByID(id string) (*realm.Entity, error) {
	r, err := s.repository.FindOneByID(id)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	if r == nil {
		return nil, nil
	}

	return r, nil
}
