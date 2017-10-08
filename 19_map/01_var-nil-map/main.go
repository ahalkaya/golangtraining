package main

import "fmt"

func main() {

	var myNilMap map[string]string
	fmt.Println(myNilMap)        // map[]
	fmt.Println(myNilMap == nil) // true

	// Declaring a nil map with var and without
	// neither _make_ nor composite literal
	// this is not the way we want to use

	// We cannot use _append_ with a map

	myNilMap["k1"] = "Tom"   // panic: assignment to entry in nil map
	myNilMap["k2"] = "Jerry" // panic: assignment to entry in nil map

	// ------------------------------------------

	var myNilSlice []string
	fmt.Println(myNilSlice)        // []
	fmt.Println(myNilSlice == nil) // true

	// We can still use _append_
	myNilSlice = append(myNilSlice, "Hello")
	fmt.Println(myNilSlice)        // [Hello]
	fmt.Println(myNilSlice == nil) // false

	myNilSlice[0] = "No way"       // runtime error: index out of range
	myNilSlice[1] = "Just kidding" // runtime error: index out of range

}

/*
A slice is a reference type pointing an underlying array and
a map is a reference type pointing an underlying hash table.
*/
