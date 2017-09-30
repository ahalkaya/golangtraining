// Passing by value
package main

import "fmt"

func main() {
	age := 30
	changeMe(age)
	fmt.Printf("Age: %d \n", age)
}

func changeMe(z int) {
	fmt.Printf("Z: %d \n", z)
	z = 24
}

// when changeMe is called on line 8
// the value 30 is being passed as an argument
