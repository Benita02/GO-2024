package main

import "fmt"

func main() {
	//To create a channel
	ch := make(chan int)
	ch <- 5
	value := <-ch
	fmt.Println(value)

}
