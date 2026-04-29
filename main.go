package main

import (
	"fmt"
	"net/http"

	"go_practice/data"
	"go_practice/handlers"
)

func main() {
	fmt.Println("Server running on port 3000...")

	data.InitData()

	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/users/", handlers.UserByIDHandler)

	http.ListenAndServe(":3000", nil)
}





































































































































// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"
// 	"strings"
// )

// type User struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

// var users []User

// func main() {
// 	fmt.Println("Server running on port 3000...")

// 	users = []User{
// 		{ID: 1, Name: "Alice"},
// 		{ID: 2, Name: "Bob"},
// 	}

// 	http.HandleFunc("/users", usersHandler)
// 	http.HandleFunc("/users/", userByIDHandler)

// 	http.ListenAndServe(":3000", nil)
// }

// func usersHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	switch r.Method {

// 	case http.MethodGet:
// 		json.NewEncoder(w).Encode(users)

// 	case http.MethodPost:
// 		var newUser User

// 		err := json.NewDecoder(r.Body).Decode(&newUser)
// 		if err != nil {
// 			http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 			return
// 		}

// 		newUser.ID = len(users) + 1
// 		users = append(users, newUser)

// 		json.NewEncoder(w).Encode(newUser)

// 	default:
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

// // func userByIDHandler(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "application/json")

// // 	// Extract ID from URL
// // 	// Example URL: /users/2
// // 	pathParts := strings.Split(r.URL.Path, "/")

// // 	if len(pathParts) < 3 {
// // 		http.Error(w, "User ID missing", http.StatusBadRequest)
// // 		return
// // 	}

// // 	idStr := pathParts[2]

// // 	id, err := strconv.Atoi(idStr)
// // 	if err != nil {
// // 		http.Error(w, "Invalid user ID", http.StatusBadRequest)
// // 		return
// // 	}

// // 	// Search user
// // 	for _, user := range users {
// // 		if user.ID == id {
// // 			json.NewEncoder(w).Encode(user)
// // 			return
// // 		}
// // 	}

// // 	http.Error(w, "User not found", http.StatusNotFound)
// // }

// func userByIDHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// Extract ID from URL
// 	// Example: /users/2
// 	pathParts := strings.Split(r.URL.Path, "/")

// 	if len(pathParts) < 3 {
// 		http.Error(w, "User ID missing", http.StatusBadRequest)
// 		return
// 	}

// 	idStr := pathParts[2]

// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid user ID", http.StatusBadRequest)
// 		return
// 	}

// 	switch r.Method {

// 	case http.MethodGet:
// 		for _, user := range users {
// 			if user.ID == id {
// 				json.NewEncoder(w).Encode(user)
// 				return
// 			}
// 		}
// 		http.Error(w, "User not found", http.StatusNotFound)

// 	case http.MethodPut:
// 		var updatedData User

// 		// Read JSON body
// 		err := json.NewDecoder(r.Body).Decode(&updatedData)
// 		if err != nil {
// 			http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 			return
// 		}

// 		// Find and update user
// 		for i, user := range users {
// 			if user.ID == id {
// 				users[i].Name = updatedData.Name

// 				json.NewEncoder(w).Encode(users[i])
// 				return
// 			}
// 		}

// 		http.Error(w, "User not found", http.StatusNotFound)

// 	case http.MethodDelete:
// 		for i, user := range users {
// 			if user.ID == id {
// 				users = append(users[:i], users[i+1:]...)

// 				json.NewEncoder(w).Encode(map[string]string{
// 					"message": "User deleted successfully",
// 				})
// 				return
// 			}
// 		}

// 		http.Error(w, "User not found", http.StatusNotFound)

// 	default:
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }