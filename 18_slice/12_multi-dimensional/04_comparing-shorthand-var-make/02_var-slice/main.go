package main

import (
	"fmt"
)

func main() {
	var student []string
	var students [][]string

	// runtime error: index out of range
	// we cannot assign a value to an index
	student[0] = "Todd"

	// Correct way is to use append
	// student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}
