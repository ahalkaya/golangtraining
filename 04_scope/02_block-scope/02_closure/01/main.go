package main

import "fmt"

func main() {
	x := 30
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "This will live in that eclosed block."
		fmt.Println(y)
	}
	// fmt.Println(y) // outside scope of y
}
