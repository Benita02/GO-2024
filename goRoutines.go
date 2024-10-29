// THREADING AND GO ROUTINES
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	balance int
	mu      sync.Mutex
)

func credit() {
	for i := 0; i < 5; i++ {
		mu.Lock() // Lock the mutex before accessing balance
		balance += 100
		mu.Unlock() // Unlock it after updating
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println("After crediting, balance is", balance)
	}
}

func debit() {
	for i := 0; i < 5; i++ {
		mu.Lock() // Lock before accessing balance
		balance -= 100
		mu.Unlock() // Unlock after updating
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println("After debiting, balance is", balance)
	}
}

func main() {
	balance = 200
	fmt.Println("Initial balance is", balance)
	go credit()
	go debit()
	fmt.Scanln()
}
