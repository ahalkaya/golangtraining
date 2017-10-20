/*
Goroutines
https://golang.org/doc/effective_go.html#goroutines
They're called goroutines because the existing terms—threads, coroutines, processes, and so on—convey inaccurate connotations.
A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space.
It is lightweight, costing little more than the allocation of stack space. And the stacks start small, so they are cheap,
and grow by allocating (and freeing) heap storage as required.

Goroutines are multiplexed onto multiple OS threads so if one should block, such as while waiting for I/O, others continue to run.
Their design hides many of the complexities of thread creation and management.

Prefix a function or method call with the go keyword to run the call in a new goroutine.
When the call completes, the goroutine exits, silently.
(The effect is similar to the Unix shell's & notation for running a command in the background.)
*/
package main

import (
	"fmt"
)

// Here we have three goroutines.
// main(), foo() and bar().
// These two goroutines runs concurrently and
// seperately from the main function. Because of this
// we see no output in the console.
// To see that we need to use a WaitGroup.
// Which is described in the next example.
func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}
