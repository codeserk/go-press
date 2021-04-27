package service

import "press/core/user"

// GenerateForUser Generates a JWT token for the given user.
func (s *service) GenerateJWTForUser(user *user.Entity) (string ,error) {
	return s.jwt.GenerateFromUserID(user.ID.Hex())
}