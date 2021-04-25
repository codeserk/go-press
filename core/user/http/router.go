package http

import (
	"press/core/user/service"

	"github.com/gorilla/mux"
)

func MakeHandlers(router *mux.Router, userService service.Interface) {
	router.Handle("/v1/auth/register", register(userService)).Methods("POST", "OPTIONS")
}
