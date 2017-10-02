// Creating a slice using _make_ keyword with
// var notation
package main

import "fmt"

func main() {
	// make([]T, length_of_slice, capacity_of_underlying_array)
	var sx = make([]int, 0, 3)

	fmt.Println("-----------------")
	fmt.Println(sx)
	fmt.Println(len(sx))
	fmt.Println(cap(sx))
	fmt.Println("-----------------")

	for i := 0; i < 80; i++ {
		sx = append(sx, i)
		fmt.Println("Len:", len(sx), "Capacity:", cap(sx), "Value: ", sx[i])
	}

}
