package people

import "fmt"

type People struct {
	Name string
	Age  int
}

func (p People) ageLimit() {
	if p.Age < 60 {
		fmt.Println("You're accepted into the program")
	} else {
		fmt.Println("I'm sorry, but you don't meet the age requirement outlined in our guidelines")
	}
}
