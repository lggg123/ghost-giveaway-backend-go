// database/database.go
package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// Initialize the database connection logic
}

func CloseDB() {
	DB.Close()
}

func MigrateDB() {
	// Database migration logic
}
