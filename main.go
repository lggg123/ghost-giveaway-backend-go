// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lggg123/ghost-giveaway-backend-go/database"
	"github.com/rs/cors"
)

func main() {
	database.InitDB()
	defer database.CloseDB()
	database.MigrateDB()

	router := mux.NewRouter()
	router.HandleFunc("/signup", signupHandler).Methods("POST")

	corsHandler := cors.Default().Handler(router)

	// Determine the port from the environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)

	log.Printf("Server is running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), corsHandler))
}
