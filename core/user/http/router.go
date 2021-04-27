package http

import (
	"press/core/user/service"

	"github.com/gorilla/mux"
)

func MakeHandlers(router *mux.Router, userService service.Interface) {
	// Auth endpoints
	router.Handle("/v1/auth/register", register(userService)).Methods("POST", "OPTIONS")
	router.Handle("/v1/auth/login", login(userService)).Methods("POST", "OPTIONS")

	// User endpoints
	router.Handle("/v1/user/me", me(userService)).Methods("GET", "OPTIONS")
}
