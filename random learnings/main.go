package main

import (
	"fmt"
	"time"
)

func addNum(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return // you can still use "return sum"
}

func formatTime(format string) {
	fmt.Println(time.Now().Format(format))
}
func displayDate() {
	fmt.Println(time.Now().Date())
}
func addNums(total int, nums ...int) int {
	fmt.Printf("%T", nums) // what does the %T format specifier mean, do?
	for _, n := range nums {
		total += n
	}
	return total
}

// var i func() int, initializing an anonymous function
func fib() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f1, f2 = f2, f1+f2
		return f1
	}
} //WAS THERE NEED FOR AN ANONYMOUS FUNCTION IN THE FIB() FUNCTION?``
// LEARN THE filter(), map(), and reduce() functions.withclosures?

func filter(arr []int, cond func(int) bool) []int {
	result := []int{}
	for _, v := range arr {
		if cond(v) {
			result = append(result, v)
		}
	}
	return result
}
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 965, 1060}
	evens := filter(a,
		func(val int) bool {
			return val%2 == 0
		})
	fmt.Println(evens)

	gen := fib()
	for i := 0; i <= 10; i++ {
		fmt.Println(gen())
	}
	// fmt.Println(addNums(1, 2, 3, 4, 5)) // 15
	// fmt.Println(addNum(39, 100))
	// formatTime("Monday 02-01-2006 15:04 MST time")
	// displayDate()

	// Outerloop:
	// 	for i := 1; i <= 5; i++ {
	// 		for j := 1; j <= 5; j++ {
	// 			if j == 3 {
	// 				break Outerloop //The use of a label Outerloop
	// 			}
	// 			fmt.Printf("%d * %d = %d\n", i, j, i*j)
	// 		}
	// 		fmt.Println("-----------")

	// 	}

	// 	for pos, char := range "こんにちは世界" {
	// 		fmt.Printf("%c at byte location %d\n", char, pos)
	// 	}

	// 	for pos, char := range "Hey Hot Stuff!" {
	// 		fmt.Println(pos, char)
	// 		// fmt.Printf("%d, %c\n", pos, char)
	// 	}
}
