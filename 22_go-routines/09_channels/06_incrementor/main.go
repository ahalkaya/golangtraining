package main

import "fmt"

func main() {

	c1 := incrementor("Foo:")
	c2 := incrementor("Bar:")
	c4 := puller(c1)
	c5 := puller(c2)
	fmt.Println("Final Counter:", <-c4+<-c5)
}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- 1
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
