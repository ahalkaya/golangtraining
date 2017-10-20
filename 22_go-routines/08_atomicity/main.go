/*
Package atomic
https://golang.org/pkg/sync/atomic/
Package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.

These functions require great care to be used correctly. Except for special, low-level applications,
synchronization is BETTER done with CHANNELS or the facilities of the SYNC PACKAGE.
Share memory by communicating; don't communicate by sharing memory.

func AddUint64
https://golang.org/pkg/sync/atomic/#AddUint64
AddInt64 atomically adds delta to *addr and returns the new value.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var counter int64
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
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		// AddInt64 atomically adds delta to *addr and returns the new value.
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", counter)
	}
}

// go run -race main.go
// vs
// go run main.go
