package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	port       = ":5000" // Define the port as a constant
	apiBaseURL = "/api/v1/"
)

// LoggingMiddleware logs each incoming request.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r) // Call the next handler
		log.Printf("Completed in %s", time.Since(start))
	})
}

// home handles requests to the base API route.
func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK) + ": Welcome to the REST API!"))
}

func main() {
	router := mux.NewRouter()

	// Add middleware to the router
	router.Use(LoggingMiddleware)

	// Define routes with explicit HTTP methods
	router.HandleFunc(apiBaseURL, home).Methods(http.MethodGet)

	log.Printf("Server is starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// reading up on more of this again
// Learnt about middleware integration
