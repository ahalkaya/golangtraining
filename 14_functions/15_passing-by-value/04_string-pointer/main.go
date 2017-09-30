package main

import "fmt"

func main() {
	name := "Jon"
	fmt.Println(&name) // 0xc42000e1d0

	changeMe(&name)

	fmt.Println(&name) // 0xc42000e1d0
	fmt.Println(name)  // Jane

}

func changeMe(n *string) { // pointer which points an address with string value
	fmt.Println(n)  // 0xc42000e1d0
	fmt.Println(*n) // Jon
	*n = "Jane"
	fmt.Println(n)  // 0xc42000e1d0
	fmt.Println(*n) // Jane
}
