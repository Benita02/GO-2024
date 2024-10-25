package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type People struct {
	Firstname string
	Lastname  string
	/*Ensure to capitalize first letter for the fields, it won’t be exported beyond the current
	  package, and the fields won’t be visible to the json.
	  Unmarshal() function*/
}

type Name struct {
	FirstName string
	LastName  string
}
type Address struct {
	Line1 string
	Line2 string
	Line3 string
}
type Customer struct {
	Name    Name
	Email   string
	Address Address
	DOB     time.Time
}

func main() {
	//var person People

	// jsonString := `{"firstname":"Wei-Meng", "lastname":"Lee"}`

	// err := json.Unmarshal([]byte(jsonString), &person)
	// if err == nil {
	// 	fmt.Println(person.Firstname, person.Lastname)
	// } else {
	// 	fmt.Println(err)
	// }

	// MAPPING CUSTOM ATTRIBUTES
	type Rates struct {
		Base   string `json:"base currency"` //These struct field tags in back ticks are necessary due to the space between base and currency
		Symbol string `json:"destination currency"`
	}
	//MAPPING UNSTRUCTURED DATA
	jsonString := `{
						"success": true,
						"timestamp": 1588779306,
						"base": "EUR",
						"date": "2020-05-06",
						"rates": {
							"AUD": 1.683349,
							"CAD": 1.528643,
							"GBP": 0.874757,
							"SGD": 1.534513,
							"USD": 1.080054
						}
					}`
	var result map[string]interface{}

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	rates := result["rates"]
	fmt.Println(rates)
	//To get the value of SGD,
	//you first need to assert rates
	//to a map type with expected key/value types
	currencies := rates.(map[string]interface{})
	SGD := currencies["SGD"]
	fmt.Println(SGD)

	//To encode to JSON we can use json.Marshal(nameOfStruct) but this doesn't have proper indentation
	//For proper indentation use json.MarshalIndent(nameOfStruct, "", "   ")
	/*
				The second argument specifies the string to prefix to the beginning of each line of
		output, and the third argument specifies the string to indent for each line
	*/
	layoutISO := "2006-01-02"
	dob, _ := time.Parse(layoutISO, "2010-01-18")
	john := Customer{
		Name: Name{FirstName: "John",
			LastName: "Smith"},
		Email: "johnsmith@example.com",
		Address: Address{
			Line1: "The White House",
			Line2: "1600 Pennsylvania Avenue NW",
			Line3: "Washington, DC 20500",
		},
		DOB: dob,
	}
	johnJson, err := json.MarshalIndent(john, "", "/t")
	if err == nil {
		fmt.Println(string(johnJson))
	} else {
		fmt.Println(err)
	}
}
