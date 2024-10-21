/* A struct (short for structure)
is a used-defined type that allows you to
group related data into a single logical
unit. */

package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

type Point struct {
	X    float32
	Y    float32
	Z    float32
	Name []string
}

type point struct {
	x float32
	y int
	z string
} // the fields in a struct don't have to be of the same type

func newPoint(x float32, y int, z string) *point { //returns a pointer to a point struct as indicated by '*'
	p := point{x: x, y: y, z: z}
	return &p // returns address for point struct as indicated '&'
}

func main() {
	//METHODS IN GO
	// value receivers and pointer receivers,
	//do a small project to better understand

	// COMPARING STRUCTS, imported cmp package to properly compare even though fields aren't all comparable
	pt1 := Point{X: 5.6, Y: 3.8, Z: 6.9,
		Name: []string{"pt1"}}
	pt2 := Point{X: 5.6, Y: 3.8, Z: 6.9,
		Name: []string{"pt"}}
	pt3 := Point{X: 5.6, Y: 3.8, Z: 6.9,
		Name: []string{"pt"}}
	fmt.Println(cmp.Equal(pt1, pt2)) // false
	fmt.Println(cmp.Equal(pt2, pt3)) // true

	// 	var pt1 point
	// 	pt1.x = 3.1
	// 	pt1.y = 5.7
	// 	pt1.z = 4.2

	// 	// To access the value of each field individually:
	// 	fmt.Println(pt1.x)
	// 	fmt.Println(pt1.y)
	// 	fmt.Println(pt1.z)

	// 	//Using struct literals, another way to create and initialize structs
	// 	pt2 := point{x: 5.6, y: 3.8, z: 6.9}
	// 	//Field names can be omitted, You can also leave out a specific field when initializing the struct
	// 	pt3 := point{5.6, 3.8}

	// 	/*In Go, creating functions to create and initialize structs is idiomatic
	// 	(meaning it follows the style of writing Go code).
	// 	These types of functions are called constructor functions.*/

	// 	pt4 := newPoint(7.8, 9.1, 2.3)
	// 	fmt.Println(pt4) // &{7.8 9.1 2.3}

}
