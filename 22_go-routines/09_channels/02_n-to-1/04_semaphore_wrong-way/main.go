/*
Below code gives deadlock error.
First we have unbuffered channel which means that the sender blocks until
the reciever has recived the value.
Let's say the first goroutine tries to put the value i into the channel c
but we block the for loop to take the value off. The code is blocked before
the loop.
*/
package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan bool)

	// First goroutine
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	// Second goroutine
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	// we block here until done <- true
	<-done
	<-done
	close(c)

	// to unblock above
	// we need to take values off of chan c here
	// but we never get here, because we're blocked above
	for n := range c {
		fmt.Println(n)
	}

}
