package main

import (
	"encoding/json"
	"fmt"
	"go_practice/db"
	"go_practice/handlers"
	"go_practice/middleware"
	"go_practice/models"
	"go_practice/router"
	"net/http"
)

func main() {
	fmt.Println("Server running on port 3000...")

	db.Connect()

	router.GET("/users", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		middleware.Auth(handlers.UsersHandler)(w, r)
	})

	// router.GET("/users", handlers.GetUsers)
	// router.GET("/users/:id", getUserByID)

	http.ListenAndServe(":3000", http.HandlerFunc(router.Serve))
}

func createUser(w http.ResponseWriter, r *http.Request, params map[string]string) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := db.DB.QueryRow(
		"INSERT INTO users(name) VALUES($1) RETURNING id",
		user.Name,
	).Scan(&user.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"go_practice/data"
// 	"go_practice/router"

// )

// func main() {
// 	fmt.Println("Server running on port 3000...")

// 	data.InitData()

// 	// Routes
// 	router.GET("/users", getUsers)
// 	router.GET("/users/:id", getUserByID)
// 	router.POST("/users", createUser)
// 	router.DELETE("/users/:id", deleteUser)

// 	http.ListenAndServe(":3000", http.HandlerFunc(router.Serve))
// }

// // package main

// // import (
// // 	"fmt"
// // 	"net/http"

// // 	"go_practice/data"
// // 	"go_practice/routes"
// // )

// // func main() {
// // 	fmt.Println("Server running on port 3000...")

// // 	data.InitData()

// // 	// http.HandleFunc("/users",
// // 	// 	middleware.Logger(handlers.UsersHandler),
// // 	// )

// // 	// http.HandleFunc("/users/",
// // 	// 	middleware.Logger(handlers.UserByIDHandler),
// // 	// )

// // 	// Register all routes
// // 	routes.RegisterUserRoutes()

// // 	http.ListenAndServe(":3000", nil)
// // }
