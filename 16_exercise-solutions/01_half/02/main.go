/*
Write a function which takes an integer.
The function will have two returns.
The first return should be the argument divided by 2.
The second return should be a bool that let’s us know whether or not the argument was even.
For example:
	a) half(1) returns (0, false)
	b) half(2) returns (1, true)
*/
package main

import "fmt"

func main() {
	h, even := half(5)
	fmt.Println(h, even)
}

func half(i int) (float64, bool) {
	return float64(i) / 2, i%2 == 0
}
