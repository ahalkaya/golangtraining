package main

import (
	"fmt"
	"math"
)

// The concrete type impementing the shape interface
type square struct {
	side float64
}

// The concrete type impementing the shape interface
type circle struct {
	radius float64
}

// shape interface
type shape interface {
	area() float64
}

// square implements the shape interface
func (s square) area() float64 {
	return s.side * s.side
}

// circle implements the shape interface
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {

	s := square{10}
	c := circle{5}
	// The concrete type impementing the shape interface is square
	info(s)
	// The concrete type impementing the shape interface is circle
	info(c)

	/*
		{10}
		100
		{5}
		78.53981633974483
	*/

}
