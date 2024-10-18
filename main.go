package main

import (
	"fmt"
)

func main() {

Outerloop:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j == 3 {
				break Outerloop //The use of a label Outerloop
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("-----------")

	}

	for pos, char := range "こんにちは世界" {
		fmt.Printf("%c at byte location %d\n", char, pos)
	}

	for pos, char := range "Hey Hot Stuff!" {
		fmt.Println(pos, char)
		// fmt.Printf("%d, %c\n", pos, char)
	}
}
