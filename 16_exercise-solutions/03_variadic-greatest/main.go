/*
Write a function with one variadic parameter that
finds the greatest number in a list of numbers.
*/
package main

import "fmt"

func main() {
	fmt.Println(max(19, 5, 3, 4, 21, 24, 36, 9, 7, 11, 92))
	fmt.Println(max(-200, -500, -10))
}

func max(numbers ...int) int {
	var largest int
	// Since zero is greater than any negative number
	// we check the index value to assign the first
	// value of _largest_ which is originally zero
	for i, v := range numbers {
		if v > largest || i == 0 {
			largest = v
		}
	}
	return largest
}
