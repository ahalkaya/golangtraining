/*
Channels
Like maps, channels are allocated with make, and the resulting value acts as a reference to an underlying data structure.
If an optional integer parameter is provided, it sets the buffer size for the channel.
The default is zero, for an unbuffered or synchronous channel.
Unbuffered channels combine communication—the exchange of a value—with synchronization—guaranteeing that
two calculations (goroutines) are in a known state.

 - RECEIVERS ALWAYS BLOCK UNTIL THERE IS DATA TO RECEIVE.
 - If the channel is unbuffered, the SENDER BLOCKS UNTIL RECEIVER HAS RECIVED the value.
 - If the channel has a buffer, the sender blocks only until the value has been copied to the buffer;
 	- if the buffer is full, this means waiting until some receiver has retrieved a value.
 - A buffered channel can be used like a semaphore, for instance to limit throughput.
 - It’s possible to CLOSE a NON-EMPTY channel but STILL have the remaining VALUES be RECIVED.
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			// sender blocks until receiver has recieved the value.
			c <- i
		}
	}()

	go func() {
		for {
			// receivers always block until there is data to receive.
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)

}
