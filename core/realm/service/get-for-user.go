package service

import (
	"fmt"
	"press/core/realm"
	"press/core/user"
)

func (s *service) GetForUser(author *user.Entity) ([]*realm.Entity, error) {
	result, err := s.repository.FindByAuthor(author)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	if result == nil {
		return nil, nil
	}

	return result, nil
}