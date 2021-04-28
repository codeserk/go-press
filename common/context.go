package common

import (
	"context"
	"press/common/constants"
	"press/core/user"
)

// GetUser Gets the user from the context.
func GetUser(ctx context.Context) *user.Entity {
	if userInContext, ok := ctx.Value(constants.UserContextKey).(*user.Entity); ok {
		return userInContext
	}

	return nil
}
