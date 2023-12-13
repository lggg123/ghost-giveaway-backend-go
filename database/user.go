// user.go (or any appropriate file)
package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Add other fields as needed
}
