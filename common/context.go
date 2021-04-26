package common

import (
	"context"
	"press/common/constants"
	"press/core/user"
)

// GetUser Gets the user from the context
func GetUser(context context.Context) (*user.Entity) {
	if user, ok := context.Value(constants.UserContextKey).(*user.Entity); ok {
		return user
	}

	return nil
}