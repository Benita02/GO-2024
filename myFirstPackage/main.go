package main

import (
	"fmt"
	"myFirstPackage/geometry"
)

func main() {
	pt1 := geometry.Point{X: 2, Y: 3}
	fmt.Println(pt1)
	fmt.Println(pt1.Length())
}
