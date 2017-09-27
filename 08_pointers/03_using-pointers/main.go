package main

import "fmt"

func main() {

	a := 30

	fmt.Println(a)  // 30
	fmt.Println(&a) // 0xc420012088

	var b *int = &a
	fmt.Println(b)  // 0xc420012088
	fmt.Println(*b) // 30

	*b = 5         // b says, "The value at this address, change it to 5"
	fmt.Println(a) // 5

	// this is useful
	// we can pass a memory address instead of a bunch of values (we can pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around addresses

	// everything is PASS BY VALUE in go, btw
	// when we pass a memory address, we are passing a value
}
