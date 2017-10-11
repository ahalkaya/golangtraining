package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []string{"Zeno", "John", "Al", "Jenny"}
	// sort.Strings(s)
	// sort.StringSlice(s).Sort()
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

}

// StringSlice attaches the methods of Interface to []string, sorting in increasing order.
