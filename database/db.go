// database/db.go
package database

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {

	host := "postgres"
	port := 5432
	user := "postgres"
	password := "password"
	dbname := "ghosterz"
	// Connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	// Connect to the database
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database")
}

// CloseDB closes the database connection
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}

func MigrateDB() {
	// Perform your database migrations here
	models := []interface{}{
		&User{},
		// Add other models as needed
	}

	// AutoMigrate all registered models
	if err := DB.AutoMigrate(models...); err != nil {
		log.Fatal(err)
	}
	log.Println("Database migration completed!")
}
