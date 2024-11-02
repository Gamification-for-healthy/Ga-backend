package main

import (
	"Ga-backend/routes"
	"log"
	"net/http"
	"os"
)

func main() {
    // Define the port to listen on
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default to port 8080 if PORT environment variable is not set
    }

    // Set up routes and start the server
    router := routes.URL() // Assuming routes.URL() returns an http.Handler or *mux.Router

    log.Printf("Starting server on port %s...", port)
    if err := http.ListenAndServe(":"+port, router); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
