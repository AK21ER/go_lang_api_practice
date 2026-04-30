package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := "host=localhost port=5432 user=postgres password= dbname=goapi sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB not reachable:", err)
	}

	log.Println("Connected to PostgreSQL ✅")
}