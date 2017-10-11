package main

import (
	"fmt"
)

// square is a user defined type and
// the underlying type is a struct.
type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

// shape is a user defined type and the underlying type is an interface.
// If we want to implement shape interface we need to have a method
// whit this signature. Needs to be called "area", needs have no
// parameters and needs to return "float64"
type shape interface {
	// method signature name and return type
	area() float64
}

func info(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {

	s := square{10}
	fmt.Printf("%T \n", s)
	info(s)

	/*
		main.square
		{10}
		100
	*/

}
