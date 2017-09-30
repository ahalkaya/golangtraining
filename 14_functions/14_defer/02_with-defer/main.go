// Example with defer - deferred function
package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func sup() {
	fmt.Println("What's up?")
}

func main() {
	// Deferred functions are executed in LIFO order
	defer sup()   // the last executed
	defer world() // before the last executed
	hello()
}

/*
Go's defer statement schedules a function call (the deferred function)
to be run immediately before the function executing the defer (in that case it is the main func) returns.
*/
