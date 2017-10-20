// 1 sender 2 reciever
package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	// Receiver 1
	go func() {
		for n := range c {
			fmt.Printf("%d - Reciever 1 \n", n)
		}
		done <- true
	}()

	// Receiver 2
	go func() {
		for n := range c {
			fmt.Printf("%d - Reciever 2 \n", n)
		}
		done <- true
	}()

	<-done
	<-done
}
