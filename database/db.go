// database/db.go
package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "ghosterz"
)

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the PostgreSQL database")
}

func CloseDB() {
	db.Close()
}

func MigrateDB() {
	// Implement your migration logic here
	// Example: db.Exec("CREATE TABLE IF NOT EXISTS users (id serial primary key, username varchar, email varchar);")
}
