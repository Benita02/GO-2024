// Working with third party packages
package main

import (
	"fmt"

	"github.com/hackebrot/turtle"
)

func main() {
	emoji, ok := turtle.Emojis["smiley"]
	if !ok {
		fmt.Println("No emoji found.")
	} else {
		fmt.Println(emoji.Char)
	}
}

// moving on to work on go doc package which
