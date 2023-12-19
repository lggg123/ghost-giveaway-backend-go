// handlers.go
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lggg123/ghost-giveaway-backend-go/database"
	"github.com/lggg123/ghost-giveaway-backend-go/models"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the user to the database
	if err := database.DB.Create(&user).Error; err != nil {
		log.Println("Error creating user in the database:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
