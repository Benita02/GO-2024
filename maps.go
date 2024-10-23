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

type kv struct {
	key   string
	value int
}

// use sort.Strings(keys) to sort keys alphabeltically from sort package
type kvPairs []kv
type dob struct {
	day   int
	month int
	year  int
}
type people struct {
	name  string
	email string
	dob   dob
}

func main() {
	// USING STRUCTS AND MAPS IN GO
	members := make(map[int]people)

	members[1] = people{
		name:  "Vex'ahlia Griffin ",
		email: "vox-machina@amazon.com",
		dob: dob{
			day:   12,
			month: 10,
			year:  1996,
		},
	}
	members[2] = people{
		name:  "John Smith",
		email: "johnsmith@example.com",
		dob: dob{
			day:   9,
			month: 12,
			year:  1988,
		},
	}
	members[3] = people{
		name:  "Janet Doe",
		email: "janetdoe@example.com",
		dob: dob{
			day:   1,
			month: 12,
			year:  1988,
		},
	}
	members[4] = people{
		name:  "Adam Jones",
		email: "adamjones@example.com",
		dob: dob{
			day:   19,
			month: 8,
			year:  2001,
		},
	}
	for k, v := range members {
		fmt.Println(k, v.name, v.email, v.dob)
	}

	//SORTING A MAP OF STRUCTS

	// lengths := make(map[string]int)

	// lengths["Toyin_land"] = 750

	// // Checking the existence of a key
	// if v, ok := heights["Jer"]; ok {
	// 	fmt.Println("Height of Jer:", v)
	// } else {
	// 	fmt.Println("Key 'Jer' does not exist!")
	// }

	// // Use delete() to remove a key from the map
	// if v, ok := heights["Bitch"]; ok {
	// 	delete(heights, "Bitch")
	// 	fmt.Println("Deleted 'Bitch' with height:", v)
	// } else {
	// 	fmt.Println("Key 'Bitch' does not exist!")
	// }

	// for k, v := range heights {
	// 	fmt.Println(k, v)
	// }

	p := make(kvPairs, len(heights))
	i := 0
	for k, v := range heights {
		p[i] = kv{k, v}

		i++
	}
	fmt.Println(p)
}
