// THREADING AND GO ROUTINES
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	balance int64      //changed type to int64 bcos of use of atomic counters
	mu      sync.Mutex //understand better when to use &sync.Mutex{} instead
)

// func credit() {
// 	for i := 0; i < 5; i++ {
// 		mu.Lock() // Lock the mutex before accessing balance
// 		balance += 100
// 		mu.Unlock() // Unlock it after updating
// 		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
// 		fmt.Println("After crediting, balance is", balance)
// 	}
// }

// func debit() {
// 	for i := 0; i < 5; i++ {
// 		mu.Lock() // Lock before accessing balance
// 		balance -= 100
// 		mu.Unlock() // Unlock after updating
// 		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
// 		fmt.Println("After debiting, balance is", balance)
// 	}
// }

// USING ATOMIC COUNTERS
func credit() {
	for i := 0; i < 10; i++ {
		// adds 100 to balance atomically
		atomic.AddInt64(&balance, 100)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}
func debit() {
	for i := 0; i < 5; i++ {
		// deducts -100 from balance atomically
		atomic.AddInt64(&balance, -100)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}
func main() {
	balance = 200
	fmt.Println("Initial balance is", balance)
	go credit()
	go debit()
	fmt.Scanln()
	fmt.Println("Final balance:", balance)
}
