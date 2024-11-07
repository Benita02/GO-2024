// Working with third party packages
package main

import (
	"fmt"

	"github.com/hackebrot/turtle"
)

func main() {
	emoji, ok := turtle.Emojis["middle finger"]
	if !ok {
		fmt.Println("No emoji found.")
	} else {
		fmt.Println("Fuck you bitches", emoji.Char)
	}
}
