package main

import "fmt"

func main() {

	var myArray [6]int
	fmt.Printf("%T \n", myArray) // [6]int
	fmt.Println(myArray)         // [0 0 0 0 0 0]

	var myEmptySlice []int
	fmt.Printf("%T \n", myEmptySlice)                 // []int
	fmt.Println(myEmptySlice)                         // []
	fmt.Println(len(myEmptySlice), cap(myEmptySlice)) // 0 0
	fmt.Println(myEmptySlice == nil)                  // true <<<

	// Below, with this shorthand style
	// there is an non-nil slice with 0 length and capacity.
	// Therefore it can be said that this is a shorthand
	// for make keyword to create a slice with 0 len and cap.
	mySecondSlice := []int{}
	fmt.Printf("%T \n", mySecondSlice)                  // []int
	fmt.Println(mySecondSlice)                          // []
	fmt.Println(len(mySecondSlice), cap(mySecondSlice)) // 0 0
	fmt.Println(mySecondSlice == nil)                   // false <<<

	mySlice := []int{1, 3, 5, 7, 9}
	fmt.Printf("%T \n", mySlice)            // []int
	fmt.Println(mySlice)                    // [1 3 5 7 9]
	fmt.Println(len(mySlice), cap(mySlice)) // 5 5

}

/*
A slice is a data structure that we use for storing a list of values
of a certain type.

Slices are built up of three elements:
	- A pointer to an underlying array.
	- Length: Number of elements that the slice is curently holding.
	- Capacity: Number of elements available in the underlying array.
	buid on top of an array.

A slice is a reference type.
The value of an uninitialized slice is nil. Because it is a reference type.
Slices are dynamic unlike arrays.
*/
