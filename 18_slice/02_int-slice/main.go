package main

import "fmt"

func main() {

	// nil slice with 0 length and capacity
	var sx []int
	fmt.Println(sx)               // []
	fmt.Println(sx == nil)        // true
	fmt.Println(len(sx), cap(sx)) // 0 0

	// non-nil slice with 0 length and capacity
	sy := []int{}
	fmt.Println(sy)               // []
	fmt.Println(sy == nil)        // false
	fmt.Println(len(sy), cap(sy)) // 0 0

	// non-nil slice with length 6, and capacity 6
	sz := []int{1, 3, 5, 7, 9, 11}
	fmt.Println(sz)               // [1 3 5 7 9 11]
	fmt.Println(sz == nil)        // false
	fmt.Println(len(sz), cap(sz)) // 6 6

	// loop using range
	for i, v := range sz {
		fmt.Println(i, " - ", v)
	}

}
