package main

import "fmt"

func max(x int) int {
	return 30 + x
}

func main() {
	fmt.Println(max)
	max := max(10)
	fmt.Println(max) // max is now the result, not the function
}

// don't do this; bad coding practice
