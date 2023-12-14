// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/lggg123/ghost-giveaway-backend-go/database"
	// Update with the correct module path
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()
	defer database.CloseDB()
	database.MigrateDB()

	router := mux.NewRouter()
	router.HandleFunc("/signup", signupHandler).Methods("POST")

	corsHandler := cors.Default().Handler(router)

	port := 8080
	log.Printf("Server is running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), corsHandler))
}
