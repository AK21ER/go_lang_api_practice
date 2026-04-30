package main

import (
	"fmt"
	"net/http"

	"go_practice/data"
	"go_practice/router"
	
)

func main() {
	fmt.Println("Server running on port 3000...")

	data.InitData()

	// Routes
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", createUser)
	router.DELETE("/users/:id", deleteUser)

	http.ListenAndServe(":3000", http.HandlerFunc(router.Serve))
}


func getUsers(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Fprintf(w, "All users")
}

func getUserByID(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Fprintf(w, "User ID: %s", params["id"])
}

func createUser(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Fprintf(w, "Create user")
}

func deleteUser(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Fprintf(w, "Delete user %s", params["id"])
}















// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"go_practice/data"
// 	"go_practice/routes"
// )

// func main() {
// 	fmt.Println("Server running on port 3000...")

// 	data.InitData()

// 	// http.HandleFunc("/users",
// 	// 	middleware.Logger(handlers.UsersHandler),
// 	// )

// 	// http.HandleFunc("/users/",
// 	// 	middleware.Logger(handlers.UserByIDHandler),
// 	// )

// 	// Register all routes
// 	routes.RegisterUserRoutes()

 
// 	http.ListenAndServe(":3000", nil)
// }