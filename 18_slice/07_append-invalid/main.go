package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"

	// Below statement gives a runtime error: index out of range.
	// It exceeds the slice's length.
	// To add more element to the slice without getting out of range error
	// we need to use _append_ keyword
	greeting[3] = "suprabadham"

	fmt.Println(greeting[2])
}
