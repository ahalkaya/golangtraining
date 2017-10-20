/*
A sender can close a channel to indicate that no more values will be sent.
Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression:

v, ok := <-ch
ok is false if there are no more values to receive and the channel is closed.

The loop for i := range c receives values from the channel repeatedly until it is closed.

Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

Another note: Channels aren't like files; you don't usually need to close them.
Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

https://tour.golang.org/concurrency/4
*/

package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// Not closing causes dead lock. Try it by commenting below.
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}
