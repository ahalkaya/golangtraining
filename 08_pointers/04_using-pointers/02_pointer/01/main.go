package main

import "fmt"

func zero(z *int) { // receiving a pointer to an int and it becomes a new variable
	*z = 0
}

func main() {
	x := 5
	zero(&x)       // pass over by value the memory address
	fmt.Println(x) // x is 0
}
