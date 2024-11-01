package main

import "fmt"

func main() {
	//To create a channel
	ch := make(chan int)

	go func() {
		value := <-ch
		fmt.Println(value)
	}()
	ch <- 5

}
