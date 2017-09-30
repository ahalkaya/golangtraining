// Declare a function and call the function in another function
// without using a function expression
package main

import "fmt"

func greeting() {
	fmt.Println("Hello world!")
}

func main() {
	greeting()
}
