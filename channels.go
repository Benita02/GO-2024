package main

import (
	"fmt"
	"math/rand"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
func main() {
	//To create a channel
	ch := make(chan int)

	go func() {
		value := <-ch
		fmt.Println(value)
	}()
	ch <- 5

	s := []int{}
	sliceSize := 10
	for i := 0; i < sliceSize; i++ {
		s = append(s, rand.Intn(100))
	}
	c := make(chan int)
	partSize := 2
	parts := sliceSize / partSize
	i := 0
	for i < parts {
		go sum(s[i*partSize:(i+1)*partSize], c)
		i += 1
	}
	/*Because you know that you have five separate goroutines (and, therefore, a total
	  of five values to be written to the channel), you can write a loop and try to extract
	  the values in the channel:*/
	i = 0
	total := 0
	for i < parts {
		partialSum := <-c // read from channel
		fmt.Println("Partial Sum: ", partialSum)
		total += partialSum
		i += 1
	}
	fmt.Println("Total: ", total)
	fmt.Scanln()

}
