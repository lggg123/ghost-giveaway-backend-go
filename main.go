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
	router.Methods(http.MethodOptions).HandlerFunc(optionsHandler).MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return r.URL.Path == "/signup"
	})
	router.HandleFunc("/signup", signupHandler).Methods("POST")

	corsHandler := cors.Default().Handler(router)

	http.Handle("/", corsHandler)

	// Determine the port from the environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)

	log.Printf("Server is running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), corsHandler))
}

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Respond with 200 OK for OPTIONS requests
	w.WriteHeader(http.StatusOK)
}
