// Declare and call a function in another function
// using function expression and
// print the type of the assigned variable
package main

import "fmt"

func main() {
	greeting := func(name string) {
		fmt.Println("Hello,", name)
	}

	greeting("Alihan")
	fmt.Printf("%T \n", greeting) // func(string)
}
