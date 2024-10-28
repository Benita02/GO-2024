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

type Shape interface {
	Area() float64
}
type Circle struct {
	radius float64
}
type Square struct {
	length float64
}

// Circle implements Shape
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// Square implements Shape
func (s Square) Area() float64 {
	return s.length * s.length
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
func calculateArea(listOfShapes []Shape) {
	for _, shape := range listOfShapes {
		fmt.Println("Area: ", shape.Area())
	}
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
	// HOW I MAY USE INTERFACES
	c1 := Circle{radius: 5}
	s1 := Square{length: 6}

	shapes := []Shape{c1, s1}
	calculateArea(shapes)
	/*
	   Suppose you now have a new shape, Triangle, that you want to add to your program. In this case, you just need to define the struct for the Triangle shape and
	   implement the Shape interface by providing the implementation for the Area()
	   function:*/
	t1 := Triangle{base: 6, height: 8}
	shapes = []Shape{c1, s1, t1}
	calculateArea(shapes)

	//Stringer interface in fmt.Println, check go docs to better understand
	//IMPLEMENTING MULTIPLE INTERFACES
	var v interface{} = c1
	v, ok := v.(Shape)
	fmt.Println(v, ok)

}
