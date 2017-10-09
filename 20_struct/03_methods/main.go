package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// Any value that is of the type person can call this method
func (p person) fullName() string {
	return p.first + " " + p.last
}

func main() {
	p1 := person{"James", "Bond", 29}
	p2 := person{"Lara", "Croft", 27}
	p3 := person{"Obi wan", "Kenobi", 42}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
	fmt.Println(person.fullName(p3)) // this is valid and the same as above
}
