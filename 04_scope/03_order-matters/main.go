package main

import "fmt"

func main() {
	fmt.Println(x)
	fmt.Println(y) // prints 25
	x := 30
}

var y = 25 // package level and the order does not matter
