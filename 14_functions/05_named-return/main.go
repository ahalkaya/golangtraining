// Declare a function which uses named return
package main

import "fmt"

func main() {
	fullName, length := greet("Jane", "Doe")
	fmt.Printf("Hello %s, lengt of your full name is %d \n", fullName, length)
}

func greet(fname string, lname string) (s string, l int) {
	s = fname + " " + lname
	l = len(s)
	return
}
