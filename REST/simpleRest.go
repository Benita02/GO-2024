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

    // Check if the Content-Type is application/json
    if r.Header.Get("Content-Type") == "application/json" {
        // POST is for creating a new course
        if r.Method == "POST" {
            // Read the string sent to the service
            var newCourse courseInfo
            reqBody, err := ioutil.ReadAll(r.Body)
            if err == nil {
                // Convert JSON to object
                json.Unmarshal(reqBody, &newCourse)

                // Validate the course title
                if newCourse.Title == "" {
                    w.WriteHeader(http.StatusUnprocessableEntity)
                    w.Write([]byte("422 - Please supply course information in JSON format"))
                    return
                }

                // Check if the course exists; add only if it doesn't
                if _, ok := courses[params["courseid"]]; !ok {
                    courses[params["courseid"]] = newCourse
                    w.WriteHeader(http.StatusCreated)
                    w.Write([]byte("201 - Course added: " + params["courseid"]))
                } else {
                    w.WriteHeader(http.StatusConflict)
                    w.Write([]byte("409 - Duplicate course ID"))
                }
            } else {
                w.WriteHeader(http.StatusUnprocessableEntity)
                w.Write([]byte("422 - Please supply course information in JSON format"))
            }
        }
    }
}

func main() {
// instantiate courses
courses = make(map[string]courseInfo)
router := mux.NewRouter()
router.HandleFunc("/api/v1/", home)
router.HandleFunc("/api/v1/courses", allcourses)
router.HandleFunc("/api/v1/courses/{courseid}", course).
Methods("GET", "PUT", "POST", "DELETE")
fmt.Println("Listening at port 5000")
log.Fatal(http.ListenAndServe(":5000", router))
}
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
