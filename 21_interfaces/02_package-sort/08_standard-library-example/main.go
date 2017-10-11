package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("Person %s, %d years old.", p.name, p.age)
}

type byAge []person

func (a byAge) Len() int {
	return len(a)
}

func (a byAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func (a byAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {

	// p := []person{
	// 	person{
	// 		name: "Vito",
	// 		age:  52,
	// 	},
	// 	person{
	// 		name: "Mikey",
	// 		age:  32,
	// 	},
	// 	person{
	// 		name: "Sonny",
	// 		age:  36,
	// 	},
	// 	person{
	// 		name: "Freddy",
	// 		age:  39,
	// 	},
	// }

	p := []person{
		{"Vito", 52},
		{"Mikey", 32},
		{"Sonny", 36},
		{"Freddy", 39},
	}
	fmt.Println(p[0])
	fmt.Println(p)
	sort.Sort(byAge(p))
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
