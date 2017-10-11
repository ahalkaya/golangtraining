package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {

	p := people{"Zeno", "John", "Al", "Jenny"}
	// Since we implement Len, Less and Swap methods
	// our type people is now type Interface.
	// Then we can use Sort because it takes Interface type as argument.
	sort.Sort(p)
	fmt.Println(p)

}

/*

type Interface interface {
        // Len is the number of elements in the collection.
        Len() int
        // Less reports whether the element with
        // index i should sort before the element with index j.
        Less(i, j int) bool
        // Swap swaps the elements with indexes i and j.
        Swap(i, j int)
}

*/
