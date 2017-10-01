package main

import "fmt"

func main() {
	var x [58]int
	fmt.Println(x)
	fmt.Println(len(x))
	x[42] = 777
	fmt.Println(x[42])
}
