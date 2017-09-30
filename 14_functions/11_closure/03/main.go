// Use a anonymous function and function expression
// as closure inside the main function
package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
Closure helps us limit the scope of variables used by multiple functions.
Without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope.

anonymous function
a function without a name

func expression
assigning a func to a variable
*/
