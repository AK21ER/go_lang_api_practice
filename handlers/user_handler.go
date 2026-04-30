package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go_practice/data"
	"go_practice/db"
	"go_practice/models"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {

	case http.MethodGet:
		json.NewEncoder(w).Encode(data.Users)// and user is like a store or like temp database that our data is stored 

	case http.MethodPost:
		var newUser models.User

		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		newUser.ID = len(data.Users) + 1
		data.Users = append(data.Users, newUser)

		json.NewEncoder(w).Encode(newUser)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func UserByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) < 3 {
		http.Error(w, "User ID missing", http.StatusBadRequest)
		return
	}

	idStr := pathParts[2]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case http.MethodGet:
		for _, user := range data.Users {
			if user.ID == id {
				json.NewEncoder(w).Encode(user)
				return
			}
		}
		http.Error(w, "User not found", http.StatusNotFound)

	case http.MethodPut:
		var updatedData models.User // here we can say that the model is used as a dto as of nestjs like what type of input we are expecting

		err := json.NewDecoder(r.Body).Decode(&updatedData)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		for i, user := range data.Users {
			if user.ID == id {
				data.Users[i].Name = updatedData.Name
				json.NewEncoder(w).Encode(data.Users[i])
				return
			}
		}

		http.Error(w, "User not found", http.StatusNotFound)

	case http.MethodDelete:
		for i, user := range data.Users {
			if user.ID == id {
				data.Users = append(data.Users[:i], data.Users[i+1:]...)

				json.NewEncoder(w).Encode(map[string]string{
					"message": "User deleted successfully",
				})
				return
			}
		}

		http.Error(w, "User not found", http.StatusNotFound)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

////////////////////////AFTER CONNECTING TO DB////////////////////////

func GetUsers(w http.ResponseWriter, r *http.Request, params map[string]string) {

	rows, err := db.DB.Query("SELECT id, name FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close() // means close the row reading ones we are done with it

	var users []models.User

	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name) // “Copy data from current database row into Go variables” and they have comed from teh sql.row 
		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)
}



func RegisterUser(w http.ResponseWriter, r *http.Request, params map[string]string) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := db.DB.QueryRow(
		"INSERT INTO users(name) VALUES($1) RETURNING id",
		user.Name,
	).Scan(&user.ID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(user)
}

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
