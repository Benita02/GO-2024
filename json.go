package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Firstname string
	Lastname  string
	/*Ensure to capitalize first letter for the fields, it won’t be exported beyond the current
	  package, and the fields won’t be visible to the json.
	  Unmarshal() function*/
}

func main() {
	var person People

	jsonString := `{"firstname":"Wei-Meng", "lastname":"Lee"}`

	err := json.Unmarshal([]byte(jsonString), &person)
	if err == nil {
		fmt.Println(person.Firstname, person.Lastname)
	} else {
		fmt.Println(err)
	}

}
