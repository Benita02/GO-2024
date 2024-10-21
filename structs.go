/* A struct (short for structure)
is a used-defined type that allows you to
group related data into a single logical
unit. */

package main

type point struct {
	x float32
	y int
	z string
} // the fields in a struct don't have to be of the same type

func main() {
	func main() {

		var pt1 point
		pt1.x = 3.1
		pt1.y = 5.7
		pt1.z = 4.2
	   }
	   // To access the value of each field individually:
	   fmt.Println(pt1.x)
	   fmt.Println(pt1.y)
	   fmt.Println(pt1.z)

	   //Using struct literals, another way to create and initialize structs
	   pt2 := point{x: 5.6, y: 3.8, z: 6.9}
	   //Field names can be omitted, You can also leave out a specific field when initializing the struct
	   pt3 := point{5.6, 3.8}
}
