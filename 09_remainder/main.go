package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is an even number\n", i)
		} else {
			fmt.Printf("%d is an odd number\n", i)
		}
	}
}
