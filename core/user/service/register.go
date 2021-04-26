package service

import (
	"fmt"
	"press/common/util"
	"press/core/errors"
	"press/core/user"
	"press/core/user/repository"
)

// Registers a new user with the given params
func (s *service) Register(params RegisterParams) (*user.Entity, string, error) {
	// Tries to find the user by the email
	userWithEmail, err := s.repository.FindOneByEmail(params.Email)
	if err != nil {
		return nil, "", fmt.Errorf("error found while trying to get user by email: %v", err)
	}
	if userWithEmail != nil {
		return nil, "", errors.Publicf("Email '%v' is already used", params.Email)
	}

	// Create and save user in DB.
	userToCreate := repository.CreateOneParams(params)
	userToCreate.Password, err = util.HashAndSalt(userToCreate.Password)
	if err != nil {
		return nil, "", fmt.Errorf("error while generating the hash for the password: %v", err)
	}
	createdUser, err := s.repository.CreateOne(userToCreate)
	if err != nil {
		return nil, "", fmt.Errorf("error found while trying to create a new user: %v", err)
	}

	// Generate token
	token, err := s.jwt.GenerateFromUserID(createdUser.ID.Hex())
	if err != nil {
		return nil, "", fmt.Errorf("error found while generating JWT from user id: %v", err)
	}

	return createdUser, token, nil
}
