// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"database" // Update with the correct module path

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Initialize the database connection
	database.InitDB()
	defer database.CloseDB()

	// Perform database migration
	database.MigrateDB()

	// Create a new router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/signup", signupHandler).Methods("POST")

	// CORS middleware
	corsHandler := cors.Default().Handler(router)

	// Start the server
	port := 8080
	log.Printf("Server is running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), corsHandler))
}
