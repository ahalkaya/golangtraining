package main

import (
	"fmt"
	"sort"
)

type people []string

// Len is the number of elements in the collection.
func (p people) Len() int {
	return len(p)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

// Swap swaps the elements with indexes i and j.
func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {

	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	// Since we implement Len, Less and Swap methods
	// our type people is now type Interface.
	// Then we can use Sort because it takes Interface type as argument.
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

}
