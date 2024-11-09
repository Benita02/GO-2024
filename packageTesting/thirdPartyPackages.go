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

/*moving on to work with the godoc package which is a documentation generator for Go programs. By default, Godoc looks at the packages installed on your computer by
examining the values in the $GOROOT and $GOPATH environment variables. It then
runs a web server and presents the documentation it generated as web pages.*/

//no electricity, worked from phone
