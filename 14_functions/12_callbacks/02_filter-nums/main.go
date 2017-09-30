// Define a function which takes a callback as an argument.
// The callback filters the number passing in it.
package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	// xs := []int{}
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	xs := filter([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(xs) // [0 2 4 6 8]
}
