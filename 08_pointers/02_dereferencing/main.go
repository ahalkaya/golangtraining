package main

import (
	"fmt"
)

func main() {
	a := 30

	fmt.Println(a)  // 30
	fmt.Println(&a) // 0xc420012088

	var b *int = &a
	fmt.Println(b)  // 0xc420012088
	fmt.Println(*b) // 30

	// b is an integer pointer;
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add a * in front of b
	// this is known as dereferencing
	// the * is an operator in this case
}
