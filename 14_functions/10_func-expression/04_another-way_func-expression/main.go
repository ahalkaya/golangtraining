// Declare a function which returns a function
// and assign the function to a varible - function expression
package main

import "fmt"

// returns a function which returns a string
func makeGreeter() func() string {
	return func() string {
		return "Hello world!"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
}
