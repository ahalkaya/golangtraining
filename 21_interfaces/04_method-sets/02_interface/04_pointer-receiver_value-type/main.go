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
	// It is not valid to call a pointer receiver by passing in a value.
	info(c)
	// cannot use c (type circle) as type shape in argument to info:
	// circle does not implement shape (area method has pointer receiver)

	/*

		If we have a value reciver and we pass in an address, Go will figure it out and
		dereference this address and then attach it to the method.

		But we have a pointer reciver and we pass in value it won't work.

		Reason:
		Go has a parallel type system. There TYPED and UNTYPED constants.
		UNTYPED constant is a constant value that does not yet have a fixed type is a "kind".
		It is not yet forced to obey the strict rules that prevent combining differently typed values.
		It is this notion of an untyped constant that makes it possible for us to use constants in Go
		with great freedom.
		If we did not have untyped constants, then we would have to do conversion on every literal value we used.

		In our case above example shows that we have no literal value. And that could be a type that is unknown.
		We could be passing in some type which is not yet known.
		It cannot take an address of something which does not have a type yet.
		These values can be existing anywhere. They might not even be on the stack or the heap.
		They might just be a value that is being used right in like the registers closer to the CPU.

		*When we have a value it might not have an address. The address might not exist for that value.

		*If we use a pointer reciever we have got to use a pointer in order to attach the method.
		*If we use a value reciver we could use either a pointer or a value to go in.

		Receivers       Values
		-----------------------------------------------
		(t T)           T and *T
		(t *T)          *T

		Values          Receivers
		-----------------------------------------------
		T               (t T)
		*T              (t T) and (t *T)

		the above rule applies for only interfaces...

		for all other than interfaces, the below rule applies

		Values          Receivers
		-----------------------------------------------
		T               (t T) and (t *T)
		*T              (t T) and (t *T)


	*/

}
