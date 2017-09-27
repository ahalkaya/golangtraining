package main

import "fmt"

const p string = "death & taxes"

func main() {
	const q = 30

	p := 5
	// q := "something" // cannot assign to q - not works

	// type and value are changed
	// because of the function scope - this causes varible shadowing
	fmt.Println("p - ", p) // prints 5
	fmt.Println("q - ", q)
}

// a constant is a simple unchanging value
