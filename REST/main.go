package main

import (
	"fmt"
)

// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gorilla/mux"
// )

// const (
// 	port       = ":5000" // Define the port as a constant
// 	apiBaseURL = "/api/v1/public"
// )

// // LoggingMiddleware logs each incoming request.
// func LoggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		log.Printf("Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
// 		next.ServeHTTP(w, r) // Call the next handler
// 		log.Printf("Completed in %s", time.Since(start))
// 	})
// }

// // home handles requests to the base API route.

// func main() {
// 	//You can wrap specific routes with middleware by using the Use method on a subrouter.
// 	router := mux.NewRouter()

// 	// Public route (no middleware)
// 	router.HandleFunc(apiBaseURL, func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte(http.StatusText(http.StatusOK) + ": Welcome to the REST API!"))
// 	}).Methods(http.MethodGet)

// 	// Create a subrouter for protected routes
// 	protected := router.PathPrefix("/api/v1/protected").Subrouter()
// 	// Add middleware to the router
// 	protected.Use(LoggingMiddleware) // Apply middleware only to protected routes

// 	protected.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Protected route: middleware applied"))
// 	}).Methods(http.MethodGet)

// 	log.Printf("Server is starting on port %s\n", port)
// 	log.Fatal(http.ListenAndServe(port, router))
// }

// //break
