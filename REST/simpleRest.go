package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the REST API!")
}

type courseInfo struct {
	Title string `json:"Title"`
}

var courses map[string]courseInfo

func allcourses(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "List of all courses")
	kv := r.URL.Query()
	for k, v := range kv {
		fmt.Println(k, v)
	}
	// returns all the courses in JSON
	json.NewEncoder(w).Encode(courses)
	// if val, ok := kv["country"]; ok {
	// 	fmt.Println(val[0])
	// }
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

//Reading up on creating new courses for this particular code
