/*
A race condition occurs when two or more threads can access shared data and they try to change it at the same time.
Because the thread scheduling algorithm can swap between threads at any time,
you don't know the order in which the threads will attempt to access the shared data.
Therefore, the result of the change in data is dependent on the thread scheduling algorithm,
i.e. both threads are "racing" to access/change the data.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var counter int
var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()

}

func incrementor(s string) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
}

// go run -race main.go
// vs
// go run main.go
