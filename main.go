package main

import (
    "log"
    "net/http"
    "os"
    "Ga-backend/routes"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    router := routes.URL()

    log.Printf("Starting server on port %s...", port)
    if err := http.ListenAndServe(":"+port, router); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}

