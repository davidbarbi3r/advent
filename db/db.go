package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

func Connect() {
	dbURL := os.Getenv("TURSO_DATABASE_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")
	url := "libsql://" + dbURL + "?authToken=" + authToken

	var err error
	DB, err = sql.Open("libsql", url)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	log.Println("Database connection established")
}
