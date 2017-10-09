package main

import "fmt"

// type: we are declaring a type
// person: name of the identifier
// struct: the underlying type is struct
type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{"James", "Bond", 29}
	p2 := person{"Lara", "Croft", 27}
	p3 := person{
		first: "Obi wan",
		last:  "Kenobi",
		age:   42,
	}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(p3.first, p3.last, p3.age)

}
