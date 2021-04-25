package service

import (
	"fmt"
	"press/core/errors"
	"press/core/user"
	"press/core/user/repository"
)

func (s *service) Register(params RegisterParams) (*user.Entity, error) {
	// Tries to find the user by the email
	userWithEmail, err := s.repository.FindOneByEmail(params.Email)
	if err != nil {
		return nil, fmt.Errorf("Error found while trying to get user by email: %v", err)
	}
	if userWithEmail != nil {
		return nil, errors.Publicf("Email '%v' is already used", params.Email)
	}

	createdUser, err := s.repository.CreateOne(repository.CreateOneParams(params))
	if err != nil {
		return nil, fmt.Errorf("Error found while trying to create a new user: %v", err)
	}

	return createdUser, nil
}
