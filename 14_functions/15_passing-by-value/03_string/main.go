package main

import "fmt"

func main() {

	name := "Jon"
	fmt.Println(name) // Jon

	changeMe(name)

	fmt.Println(name) // Jon
}

func changeMe(z string) {
	fmt.Println(z) // Jon
	z = "Jane"
	fmt.Println(z) // Jane
}
