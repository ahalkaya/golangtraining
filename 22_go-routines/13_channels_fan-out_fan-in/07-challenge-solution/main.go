package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	in := generate()

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is closed
	// Distribute work across multiple functions (ten goroutines) that all read from in.
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)
	c6 := factorial(in)
	c7 := factorial(in)
	c8 := factorial(in)
	c9 := factorial(in)
	c10 := factorial(in)

	// FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 through c9 onto a single channel
	var x int
	for n := range merge(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10) {
		x++
		fmt.Println(x, "\t", n)
	}

}

func generate() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			out <- random()
		}
		close(out)
	}()
	return out
}

func factorial(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func merge(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(channels))

	for _, ch := range channels {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func random() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func fact(n int) int {
	total := 1
	for i := n; i > 1; i-- {
		total *= i
	}
	return total
}
