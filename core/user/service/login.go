package service

import (
	"fmt"
	"press/core/errors"
	"press/core/user"

	"golang.org/x/crypto/bcrypt"
)

// Logins with the credentials
func (s *service) Login(params LoginParams) (*user.Entity, error) {
	userWithEmail, err := s.repository.FindOneByEmail(params.Email)
	if err != nil {
		return nil, fmt.Errorf("Error found while trying to get user by email: %v", err)
	}
	if userWithEmail == nil {
		return nil, errors.Public("Invalid credentials")
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(userWithEmail.Password), []byte(params.Password))
	if err != nil {
		return nil, errors.Public("Invalid credentials")
	}

	return userWithEmail, nil
}
