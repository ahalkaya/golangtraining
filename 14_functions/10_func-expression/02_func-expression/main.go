// Declare and call a function in another function
// using function expression
// print the type of the assigned variable
package main

import "fmt"

func main() {
	// Declaring a function and assigning it to
	// a variable is called function expression
	greeting := func() {
		fmt.Println("Hello world!")
	}

	greeting()
	fmt.Printf("%T \n", greeting) // func()
}
