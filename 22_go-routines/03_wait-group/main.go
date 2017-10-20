/*
Package sync
https://golang.org/pkg/sync/
Package sync provides basic synchronization primitives such as mutual exclusion locks.
Other than the Once and WaitGroup types, most are intended for use by low-level library routines.
Higher-level synchronization is BETTER done via CHANNELS and COMMUNICATION.

type WaitGroup
https://golang.org/pkg/sync/#WaitGroup
A WaitGroup waits for a collection of goroutines to finish.
The main goroutine calls Add to set the number of goroutines to wait for.
Then each of the goroutines runs and calls Done when finished.
At the same time, Wait can be used to block until all goroutines have finished.
*/
package main

import (
	"fmt"
	"sync"
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
	}
}

func bar() {
	// Decrement the counter when the goroutine completes.
	defer wg.Done()
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}
