package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Function for inserting items into a slice.
func insert(orig []int, index int, value int) ([]int, error) {
	if index < 0 {
		return nil, errors.New(
			"Index cannot be less than 0")
	}
	if index >= len(orig) {
		return append(orig, value), nil //why is this nil here??
	}

	orig = append(orig[:index+1], orig[index:]...)
	orig[index] = value
	return orig, nil
}
func delete(orig []int, index int) ([]int, error) {
	if index < 0 || index > len(orig) {
		return nil, errors.New("Index out of range")
	}
	orig = append(orig[:index], orig[index+1:]...)
	return orig, nil

}

func main() {

	// TESTING DELETE FUNCTION
	t := []int{0, 100, 2, 4, 6, 8}
	t, err := delete(t, -3)
	if err == nil {
		fmt.Println(t) // [1 2 4 5]
	} else {
		fmt.Println(err)
	}
	// TESTING INSERT FUNCTION
	// t := []int{1, 2, 3, 4, 5}
	// t, err := insert(t, 2, 9)
	// if err == nil {
	// 	fmt.Println(t) // 1 2 9 3 4 5]
	// } else {
	// 	fmt.Println(err)
	// }

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
