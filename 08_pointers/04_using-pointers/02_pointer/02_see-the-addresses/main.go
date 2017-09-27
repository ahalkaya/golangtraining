package main

import "fmt"

func zero(z *int) { // receiving a pointer to an int and it becomes a new variable
	fmt.Println("Address in func zero: ", z) // 0xc42007e008
	*z = 0
}

func main() {
	x := 5
	fmt.Println("Addres in func main", &x) // 0xc42007e008
	zero(&x)
	fmt.Println(x) // x is 0
}
