package main

import "fmt"

func main() {
	// var a int // still has a memory address without a value
	a := 30

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("%d \n", &a)
}
