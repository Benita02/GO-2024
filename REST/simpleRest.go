package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the REST API!")
}
func allcourses(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "List of all courses")
	kv := r.URL.Query()
	for k, v := range kv {
		fmt.Println(k, v)
	}
	if val, ok := kv["country"]; ok {
		fmt.Println(val[0])
	}
}
func course(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprintf(w, "Detail for course "+params["courseid"])
	fmt.Fprintf(w, "/n")
	fmt.Fprintf(w, r.Method)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/courses", allcourses)
	router.HandleFunc("/api/v1/courses/{courseid}", course).Methods("GET", "PUT", "POST", "DELETE")
	fmt.Println("Listening at port 5000")
	http.ListenAndServe(":5000", router)
}

//Read a lot about REST APIs, also read about GRAPH QL
//Downloaded books on REST APIs and generally just apis
//also better understood this code and saw an optimised idiomatic version of this code
