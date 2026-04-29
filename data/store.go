package data

import "go_practice/models"

var Users []models.User

func InitData() {
	Users = []models.User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}
}