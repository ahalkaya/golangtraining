package main

import "fmt"

func zero(z int) {
	fmt.Printf("Address in func zero: %p\n", &z) // 0xc420012090
	// fmt.Println(&z)
	z = 0
}

func main() {
	x := 5
	fmt.Printf("Address in func main: %p\n", &x) // 0xc420012088
	// fmt.Println(&x)
	zero(x)
	fmt.Println("The value of x: ", x) // x is still 5
}
