 // THREADING AND GO ROUTINES
package main
import (
 "fmt"
 "math/rand"
 "time"
)

var balance int

func credit() {
 for i := 0; i < 5; i++ {
     balance += 100
     time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
     fmt.Println("After crediting, balance is", balance)
 }
}

func debit() {
 for i := 0; i < 5; i++ {
     balance -= 100
     time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
     fmt.Println("After debiting, balance is", balance)
 }
}

func main() {
 balance = 200
 fmt.Println("Initial balance is", balance)
 go credit()
 go debit()

 fmt.Scanln() // fmt.Scanln() ensures the program doesn't terminate immediately
/*
Issues with Shared Resources
When credit() and debit() run concurrently, unpredictable results occur because they both access balance without waiting for each other. This is known as a race condition. In the output, you’ll see inconsistent results like a balance of 200 after crediting when it should be 300, or other errors in sequence.*/
    




    
}

}
