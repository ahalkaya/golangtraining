package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	p1 := &person{"James", 29}
	fmt.Println(p1)         // &{James 29}; you have a pointer and the values {James 29}
	fmt.Printf("%T \n", p1) // *main.person; the pointer to main.person
	fmt.Println(p1.name)    // James; this is like *p1.name
	fmt.Println(p1.age)     // 29
}
