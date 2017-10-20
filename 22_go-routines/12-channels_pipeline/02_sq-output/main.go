package main

import (
	"fmt"
)

func main() {

	// Set up the pipeline. And consume output.
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // 16 then 81
	}

}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums { // nums is a slice of ints
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
