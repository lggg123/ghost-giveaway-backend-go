// models/user.go
package models

import (
	"gorm.io/gorm"
)

// User represents the user model
type User struct {
	gorm.Model
	Name          string `gorm:"not null"`
	Email         string `gorm:"unique;not null"`
	WalletAddress string `gorm:"unique;not null"`
	// Add other fields as needed
}
