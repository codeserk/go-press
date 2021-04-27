package service

import (
	"context"
	"net/http"
	"press/common/constants"
	"strings"
)

func (s *service) CreateAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) == 0 || !strings.Contains(token, "bearer ") {
			next.ServeHTTP(w, r)
			return
		}

		token = strings.Replace(token, "bearer ", "", 1)
		userId, err := s.jwt.GetUserIDFromToken(token)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		user, err := s.repository.FindOneById(userId)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		newRequest := r.WithContext(context.WithValue(r.Context(), constants.UserContextKey, user))
		*r = *newRequest

		next.ServeHTTP(w, r)
	})
}