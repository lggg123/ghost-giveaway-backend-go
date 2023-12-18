// database/db.go
package database

import (
	"log"
	"os"

	"github.com/lggg123/ghost-giveaway-backend-go/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"net/url"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL must be set")
	}

	connURL, parseErr := url.Parse(dbURL)
	if parseErr != nil {
		log.Fatal("Error parsing DATABASE_URL: ", parseErr)
	}

	connStr := connURL.String()

	var err error
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
		&models.User{},
		// Add other models as needed
	}

	// AutoMigrate all registered models
	if err := DB.AutoMigrate(models...); err != nil {
		log.Fatal(err)
	}
	log.Println("Database migration completed!")
}
