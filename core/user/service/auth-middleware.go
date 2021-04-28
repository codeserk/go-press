package service

import (
	"context"
	"net/http"
	"strings"

	"press/common/constants"
)

func (s *service) CreateAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" || !strings.Contains(token, "bearer ") {
			next.ServeHTTP(w, r)
			return
		}

		token = strings.Replace(token, "bearer ", "", 1)
		userID, err := s.jwt.GetUserIDFromToken(token)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		user, err := s.repository.FindOneByID(userID)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		newRequest := r.WithContext(context.WithValue(r.Context(), constants.UserContextKey, user))
		*r = *newRequest

		next.ServeHTTP(w, r)
	})
}
