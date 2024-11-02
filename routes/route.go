package routes

import (
	"Ga-backend/config"
	"net/http"

	"github.com/gorilla/mux"
)

// Define your handler functions
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Welcome to the GA Backend!"))
}

func AnotherHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("This is another route!"))
}

// URL function to set up and return a router
func URL() *mux.Router {
    router := mux.NewRouter()

    // Middleware to handle Access Control Headers
    router.Use(accessControlMiddleware)

    // Define your routes here
    router.HandleFunc("/", HomeHandler).Methods("GET")
    router.HandleFunc("/another", AnotherHandler).Methods("GET")

    return router
}

// Middleware function to handle Access Control Headers
func accessControlMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if config.SetAccessControlHeaders(w, r) {
            return
        }
        next.ServeHTTP(w, r)
    })
}
