//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//

package main

import "fmt"

// --Requirements:
//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func personName(name string) {
	return fmt.Println("Hey, ", name, "you look good MF!")
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func print() {
	return fmt.Println("This function prints out any message")
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func addThreeNumbers(no1, no2, no3 int) int {
	return no1 + no2 + no3
}

// - Write a function that returns any number
func returnsAnyNumber(num int) {
	return print.Println(num)
}

// - Write a function that returns any two numbers
func returnsAnyTwoNumber(num1, num2 int) {
	return fmt.Println(num1, num2)
}

// - Add three numbers together using any combination of the existing functions.
func functionCombination(no1, no2, no3 int) int {
	add := addThreeNumbers(no1, no2, no3)
	return fmt.Println(add)
}

// - Print the result
// - Call every function at least once
func main() {
	personName("Benita")

	returnsAnyNumber(8)
	addThreeNumbers(5, 5, 5)
	returnsAnyTwoNumber(8, 8)
	functionCombination(9, 10, 1)

}
