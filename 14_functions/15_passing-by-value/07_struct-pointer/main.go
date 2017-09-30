package main

import "fmt"

type customer struct {
	name string
	age  int
}

func main() {
	c := customer{"Alihan", 30}
	fmt.Println(&c)      // &{Alihan 30}
	fmt.Println(&c.name) // 0xc42000a060

	changeMe(&c)

	fmt.Println(c)       // {Jon 30}
	fmt.Println(&c.name) // 0xc42000a060
}

func changeMe(z *customer) {
	fmt.Println(z)       // &{Alihan 30}
	fmt.Println(&z.name) // 0xc42000a060
	// *z = customer{"Jon", 20}
	z.name = "Jon"
	fmt.Println(z)       // &{Jon 30}
	fmt.Println(&z.name) // 0xc42000a060
}
