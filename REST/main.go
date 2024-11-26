package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port       = ":5000" // Define the port as a constant
	apiBaseURL = "/api/v1/"
)

// home handles requests to the base API route.
func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK) + ": Welcome to the REST API!"))
}

func main() {
	router := mux.NewRouter()

	// Define routes with explicit HTTP methods
	router.HandleFunc(apiBaseURL, home).Methods(http.MethodGet)

	log.Printf("Server is starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
