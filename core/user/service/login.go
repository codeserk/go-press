package service

import (
	"fmt"

	"press/common/errors"
	"press/core/user"

	"golang.org/x/crypto/bcrypt"
)

// Logins with the credentials.
func (s *service) Login(params LoginParams) (*user.Entity, string, error) {
	userWithEmail, err := s.repository.FindOneByEmail(params.Email)
	if err != nil {
		return nil, "", fmt.Errorf("error found while trying to get user by email: %v", err)
	}
	if userWithEmail == nil {
		return nil, "", errors.Public("invalid credentials")
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(userWithEmail.Password), []byte(params.Password))
	if err != nil {
		return nil, "", errors.Public("invalid credentials")
	}

	// Generate token
	token, err := s.jwt.GenerateFromUserID(userWithEmail.ID.Hex())
	if err != nil {
		return nil, "", fmt.Errorf("error found while generating JWT from user id: %v", err)
	}

	return userWithEmail, token, nil
}
