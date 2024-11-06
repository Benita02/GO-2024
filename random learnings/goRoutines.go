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
	balance int64           //changed type to int64 bcos of use of atomic counters
	mu      = &sync.Mutex{} //understand better when to use &sync.Mutex{} instead
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
func credit(wg *sync.WaitGroup) {
	// notify the WaitGroup when we are done
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mu.Lock()
		// adds 100 to balance atomically
		atomic.AddInt64(&balance, 100)
		fmt.Println("After crediting, balance is", balance)
		mu.Unlock()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}
func debit(wg *sync.WaitGroup) {
	// notify the WaitGroup when we are done
	defer wg.Done()
	for i := 0; i < 5; i++ {
		mu.Lock()
		// deducts -100 from balance atomically
		atomic.AddInt64(&balance, -100)
		fmt.Println("After debiting, balance is", balance)
		mu.Unlock()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}
func main() {
	var wg sync.WaitGroup //why didn't we declare this as a global variable
	balance = 200
	fmt.Println("Initial balance is", balance)

	wg.Add(1)
	go credit(&wg) //I don't fully understand why we're passing the waitgroup address

	wg.Add(1)
	go debit(&wg)

	wg.Wait()
	fmt.Println("Final balance:", balance)
}
