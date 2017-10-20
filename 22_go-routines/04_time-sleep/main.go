/*
func Sleep
https://golang.org/pkg/time/#Sleep
Sleep pauses the current goroutine for at least the duration d.
A negative or zero duration causes Sleep to return immediately.

type Duration int64
https://golang.org/pkg/time/#Duration
A Duration represents the elapsed time between two instants as an int64 nanosecond count.
The representation limits the largest representable duration to approximately 290 years.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	// Increment the WaitGroup counter.
	wg.Add(2)
	// Launch a goroutine.
	go foo()
	// Launch a goroutine.
	go bar()
	// Wait for all goroutines to complete.
	wg.Wait()

}

func foo() {
	// Decrement the counter when the goroutine completes.
	defer wg.Done()
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		// Sleep pauses the current goroutine for at least the duration d.
		time.Sleep(3 * time.Millisecond)
	}
}

func bar() {
	// Decrement the counter when the goroutine completes.
	defer wg.Done()
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		// Sleep pauses the current goroutine for at least the duration d.
		time.Sleep(20 * time.Millisecond)
	}
}
