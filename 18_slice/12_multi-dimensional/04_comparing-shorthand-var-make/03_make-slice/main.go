package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)

	// This works since we use make and define length and capacity
	student[0] = "Todd"

	// student = append(student, "Todd")

	fmt.Println(student)
	fmt.Println(students)
}
