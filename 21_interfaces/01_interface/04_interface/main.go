package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func totalArea(shapes ...shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

func main() {

	s := square{10}
	c := circle{5}
	info(s)
	info(c)
	fmt.Println("Total Area:", totalArea(s, c))

	/*
		{10}
		100
		{5}
		78.53981633974483
		Total Area: 178.53981633974485
	*/

}
