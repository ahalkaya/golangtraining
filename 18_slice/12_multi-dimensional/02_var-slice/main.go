package main

import (
	"fmt"
)

func main() {
	var student []string
	var students [][]string
	fmt.Println(student)         // []
	fmt.Println(students)        // []
	fmt.Println(student == nil)  // true
	fmt.Println(students == nil) // true
}

/*
When using _var_ keyword, initial value of primitives is 0
however slice, map, channel are reference types.
So using var while defining a reference type sets its initial
value to _nil_. Which means there is no address pointing that
value.
*/
