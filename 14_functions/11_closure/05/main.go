// Define a wrapper function which returns function.
// Make the returned function closes over the variable inside the wrapper function.
// See the effect of closure by calling different functions returned from the wrapper function.
package main

import "fmt"

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	incrementA := wrapper()
	incrementB := wrapper()
	fmt.Println("A: ", incrementA()) // A: 1
	fmt.Println("A: ", incrementA()) // A: 2
	fmt.Println("B: ", incrementB()) // B: 1
	fmt.Println("B: ", incrementB()) // B: 2
}
