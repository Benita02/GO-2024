// THREADING AND GO ROUTINES
package main

import "fmt"

func main() {
	go say("Hello", 3)
	go say("World", 2)
	fmt.Scanln() // fmt.Scanln() ensures the program doesn't terminate immediately
}
