package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
	fmt.Println("Done and can continue to do other work")
}
func fibo(n int, c chan int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
		time.Sleep(1 * time.Second)

	}
	close(c)
}
func main() {
	c := make(chan int)

	go fibo(10, c)
	for i := range c {
		fmt.Println(i)
	}

	//Buffered Channel
	//:= make(chan int, 10)

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
		//r := rand.New(rand.NewSource(time.Now().UnixNano())) for when I want reproducible randomness
		s = append(s, rand.Intn(100)) //this will then become r.Intn(100)
	}
	cha := make(chan int, 5)
	partSize := 2
	parts := sliceSize / partSize
	i := 0
	for i < parts {
		go sum(s[i*partSize:(i+1)*partSize], cha)
		i += 1
	}
	/*Because you know that you have five separate goroutines (and, therefore, a total
	  of five values to be written to the channel), you can write a loop and try to extract
	  the values in the channel:*/
	i = 0
	total := 0

	for i < parts {
		partialSum := <-cha // read from channel
		fmt.Println("Partial Sum: ", partialSum)
		total += partialSum
		i += 1
	}
	fmt.Println("Total: ", total)
	fmt.Scanln()

}
