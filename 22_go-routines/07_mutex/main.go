/*
type Mutex
https://golang.org/pkg/sync/#Mutex
A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
A Mutex must not be copied after first use.
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
var mutex sync.Mutex

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
		// Lock locks mutex. If the lock is already in use,
		// the calling goroutine blocks until the mutex is available.
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter:", counter)
		// Unlock unlocks mutex. It is a run-time error if mutex is not locked on entry to Unlock.
		// A locked Mutex is not associated with a particular goroutine.
		// It is allowed for one goroutine to lock a Mutex and then arrange for another goroutine to unlock it.
		mutex.Unlock()
	}
}

// go run -race main.go
// vs
// go run main.go
