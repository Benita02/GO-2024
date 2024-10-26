package main

import "fmt"

type DigitsCounter interface {
	CountOddEven() (int, int)
}

type DigitString string

func (ds DigitString) CountOddEven() (int, int) {
	odds, evens := 0, 0
	for digit := range ds {
		if digit%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return odds, evens
}
func main() {
	// IMPLEMENTING AN INTERFACE
	s := DigitString("123456789")
	fmt.Println(s.CountOddEven())

	// 	Because DigitsCounter is a type in itself, you can also create a variable of type
	// DigitsCounter and then assign the variable s to it:
	var d DigitsCounter
	d = s
	fmt.Println(d.CountOddEven())
}
