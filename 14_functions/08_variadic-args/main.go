// Declare a variadic function which takes variadic parameters
// and call the function with passing in variadic arguments
package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	// variadic args
	// pulls each item out and pass it over one at the time
	n := average(data...)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
