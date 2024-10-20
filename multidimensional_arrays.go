package main

import (
	"fmt"
	"strconv"
)

func main() {
	var table [5][6]string
	for row := 0; row < 5; row++ {
		for column := 0; column < 6; column++ {
			table[row][column] = strconv.Itoa(row) + "," + strconv.Itoa(column)
			fmt.Println(strconv.Itoa(column))
		}
	}
	fmt.Println(table)

	var cube [4][3][3]string
	for row := 0; row < 4; row++ {
		for column := 0; column < 3; column++ {
			for depth := 0; depth < 3; depth++ {
				cube[row][column][depth] =
					strconv.Itoa(row) +
						strconv.Itoa(column) +
						strconv.Itoa(depth)
			}
		}
	}
	fmt.Println(cube)

	// 	slice := make([]int, 7)
	// 	fmt.Println(slice)
}
