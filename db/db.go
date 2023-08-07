package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func ConnectToPostgres() *sql.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "1234"
		dbname   = "go_store"
	)

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
