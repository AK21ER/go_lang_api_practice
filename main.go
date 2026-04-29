package main

import (
	"fmt"
	"net/http"

	"go_practice/data"
	"go_practice/handlers"
	"go_practice/middleware"	
)

func main() {
	fmt.Println("Server running on port 3000...")

	data.InitData()

	http.HandleFunc("/users",
		middleware.Logger(handlers.UsersHandler),
	)

	http.HandleFunc("/users/",
		middleware.Logger(handlers.UserByIDHandler),
	)

	http.ListenAndServe(":3000", nil)
}