package main

import (
	"fmt"
	"sort"
)

func main() {

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	// sort.Sort(sort.IntSlice(n))
	// Sort is a convenience method for type IntSlice.
	sort.IntSlice(n).Sort()
	fmt.Println(n)

}
