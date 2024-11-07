package main

import (
	"fmt"
	"packageTesting/peopleMethods.go"
)

func main() {
	var name string
	var age int
	fmt.Printf("Input your name:")

	_, err := fmt.Scanf("%s\n", &name)
	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	fmt.Printf("Input your age:")

	_, err = fmt.Scanf("%d\n", &age)
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	p1 := peopleMethods.People{Name: name, Age: age}

	fmt.Println("Your name is: ", p1.Name)
	fmt.Println(p1.AgeLimit())
}
