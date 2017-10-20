/*
This example is for demonstrating race condition.
First and second goroutines have access to a shared variable "wg".
And both try to write to "wg".
In concurrent programming this is a situation that should be avoided.
In the third goroutine, "wg.Wait" could be called prematurely as wg could hit zero.
*/
package main

import (
	"fmt"
	"sync"
)

func main() {

	c := make(chan int)

	var wg sync.WaitGroup

	// First goroutine
	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	// Second goroutine
	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	// Third goroutine
	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
