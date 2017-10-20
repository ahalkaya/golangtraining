/*
Package runtime
https://golang.org/pkg/runtime/
Package runtime contains operations that interact with Go's runtime system,
such as functions to control goroutines. It also includes the low-level type information
used by the reflect package; see reflect's documentation for the programmable interface
to the run-time type system.

func GOMAXPROCS
https://golang.org/pkg/runtime/#GOMAXPROCS
The GOMAXPROCS variable limits the number of operating system threads that can execute
user-level Go code simultaneously. There is no limit to the number of threads
that can be blocked in system calls on behalf of Go code;
those do not count against the GOMAXPROCS limit.
This package's GOMAXPROCS function queries and changes the limit.

GOMAXPROCS function sets the maximum number of CPUs that can be executing simultaneously and returns the previous setting.
If n < 1, it does not change the current setting. The number of logical CPUs on the local machine can be queried with NumCPU.
This call will go away when the scheduler improves.

The init function
https://golang.org/doc/effective_go.html#init
Finally, each source file can define its own niladic init function to set up whatever state is required.
(Actually each file can have multiple init functions.)
And finally means finally: init is called after all the variable declarations in the package have evaluated
their initializers, and those are evaluated only after all the imported packages have been initialized.
Besides initializations that cannot be expressed as declarations, a common use of init functions is to
verify or repair correctness of the program state before real execution begins.
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

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
