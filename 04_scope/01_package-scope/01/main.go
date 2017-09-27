package main

import "fmt"

var x = 10

func main() {
	fmt.Println(x)
	foo()
	// Block level scope and the order matters
	y := 5 // scope of y starts here and goes to end of the block
	fmt.Println(y)
}

func foo() {
	fmt.Println(x)
	// fmt.Println(y) not works because y is undefined
}
