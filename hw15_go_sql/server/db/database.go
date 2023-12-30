package db

import (
	"database/sql"
	"fmt"
	"log"
)

var db sql.DB

func ExecuteQuery(command string) (*sql.Rows, error) {
	return db.Query(command)
}

func init() {
	db, err := sql.Open("pgx", "postgres://postgres:matrix@localhost:5432/Trade")
	if err != nil {
		log.Fatalf("Failed to load driver: %v", err)
	}
	defer db.Close()
	fmt.Println("Database connection established.")
}
