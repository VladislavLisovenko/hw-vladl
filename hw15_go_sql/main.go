package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	dsn := "postgres://postgres:matrix@localhost:5432/Trade"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to load driver: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}
	fmt.Println("Connection established!")
}
