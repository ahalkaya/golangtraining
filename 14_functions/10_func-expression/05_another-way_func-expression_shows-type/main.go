// Declare a function which returns a function
// and assign the function to a varible - function expression
// print the type of the assigned variable
package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Hello world!"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T \n", greet) // func() string
}
