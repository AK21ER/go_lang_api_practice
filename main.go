





package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"go_practice/models"

	"go_practice/router"
	"go_practice/db"
)

func main() {
	fmt.Println("Server running on port 3000...")

	db.Connect()

	router.GET("/users", getUsers)
	// router.GET("/users/:id", getUserByID)

	http.ListenAndServe(":3000", http.HandlerFunc(router.Serve))
}

func getUsers(w http.ResponseWriter, r *http.Request, params map[string]string) {

	rows, err := db.DB.Query("SELECT id, name FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)
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


// func getUsers(w http.ResponseWriter, r *http.Request, params map[string]string) {
// 	fmt.Fprintf(w, "All users")
// }

// func getUserByID(w http.ResponseWriter, r *http.Request, params map[string]string) {
// 	fmt.Fprintf(w, "User ID: %s", params["id"])
// }

// func createUser(w http.ResponseWriter, r *http.Request, params map[string]string) {
// 	fmt.Fprintf(w, "Create user")
// }

// func deleteUser(w http.ResponseWriter, r *http.Request, params map[string]string) {
// 	fmt.Fprintf(w, "Delete user %s", params["id"])
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