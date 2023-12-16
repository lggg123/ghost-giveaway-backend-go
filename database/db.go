// database/db.go
package database

import (
	"log"
	"os"

	"github.com/lggg123/ghost-giveaway-backend-go/models"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	connStr := os.Getenv("DATABASE_URL")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading .env file")
	}

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
