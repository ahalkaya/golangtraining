// Declare a function which takes a PARAMETER
// and call the function with passing in an ARGUMENT
package main

import "fmt"

func main() {
	greet("Jon")
	greet("Jane")
}

func greet(name string) {
	fmt.Println("Hello", name)
}
