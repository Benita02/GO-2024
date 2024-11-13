package main

import (
	"fmt"

	"github.com/Benita02/GO-2024.git/quotes"
	"github.com/Benita02/GO-2024.git/strings"
)

func main() {
	o, e := strings.CountOddEven("12345")
	fmt.Println(o, e) // 3 2
	fmt.Println(quotes.GetEmoji("bat"))
}

// learnt about using the cmd line with various commands and flags to better understand how to use a third party package/module and more about web API's in GO
