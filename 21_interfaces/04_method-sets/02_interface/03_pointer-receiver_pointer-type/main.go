package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {

	c := circle{5}

	// circle implements shape interface
	// and implementing with a pointer receiver.
	// It is valid to call a pointer receiver by passing a pointer value.
	info(&c)

}
