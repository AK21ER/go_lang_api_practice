package routes

import (
	"net/http"

	"go_practice/handlers"
	"go_practice/middleware"
)

func RegisterUserRoutes() {

	http.HandleFunc("/users",
		middleware.Logger(handlers.UsersHandler),
	)

	http.HandleFunc("/users/",
		middleware.Logger(handlers.UserByIDHandler),
	)
}