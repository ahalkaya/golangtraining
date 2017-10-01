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
	fmt.Println(half(1))
	fmt.Println(half(2))
}

func half(i int) (int, bool) {
	return i / 2, i%2 == 0
}
