// Declare a variable in package scope.
// Declare a function that increments the value of the variable by 1.
// Call that function to increment.
package main

import "fmt"

var x int

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
Closure helps us limit the scope of variables used by multiple functions.
Without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope.
*/
