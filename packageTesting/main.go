package main

import "fmt"

func main() {
	var name string
	fmt.Printf("Input your name:")

	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	fmt.Println("Your name is: ", name)
}
