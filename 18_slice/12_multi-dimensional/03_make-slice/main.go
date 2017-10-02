package main

import (
	"fmt"
)

func main() {
	// When we use make without capacity
	// both length and capacity are set to the given value
	student := make([]string, 35)
	students := make([][]string, 35)
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)  // false
	fmt.Println(students == nil) // false
}
