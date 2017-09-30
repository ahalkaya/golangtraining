// Declare a function which takes two parameters
// and call the function with passing two arguments
package main

import "fmt"

func main() {
	greetFullName("Jane", "Doe")
}

func greetFullName(fname string, lname string) {
	fmt.Println("Hello,", fname, lname)
}
