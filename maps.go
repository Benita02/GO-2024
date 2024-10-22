package main

import "fmt"

// INITIALIZING MAPS

var heights = map[string]int{
	"Joan":   168,
	"Benita": 169, //take note of comma at end
	"Phil":   182,
	"Julian": 178,
	"August": 184,
	"Bitch":  145,
}

// OR

func main() {
	lengths := make(map[string]int)

	lengths["Toyin_land"] = 750

	// Checking the existence of a key
	if v, ok := heights["Jer"]; ok {
		fmt.Println("Height of Jer:", v)
	} else {
		fmt.Println("Key 'Jer' does not exist!")
	}

	// Use delete() to remove a key from the map
	if v, ok := heights["Bitch"]; ok {
		delete(heights, "Bitch")
		fmt.Println("Deleted 'Bitch' with height:", v)
	} else {
		fmt.Println("Key 'Bitch' does not exist!")
	}

	for k, v := range heights {
		fmt.Println(k, v)
	}
}
