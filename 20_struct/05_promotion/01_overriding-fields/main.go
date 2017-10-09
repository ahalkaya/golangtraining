package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type doubleZero struct {
	person
	first         string
	licenseToKill bool
}

func main() {

	p1 := doubleZero{
		person: person{
			first: "James",
			last:  "Bond",
			age:   29,
		},
		first:         "Double Zero Seven",
		licenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			first: "Miss",
			last:  "MoneyPenny",
			age:   25,
		},
		first:         "If looks could kill",
		licenseToKill: false,
	}

	p3 := doubleZero{
		first: "Mr Noone",
	}

	// fields and methods of the inner-type are promoted to the outer-type
	fmt.Println(p1.first, p1.person.first)
	fmt.Println(p2.first, p2.person.first)
	fmt.Println(p3.first, p3.person.first) // zero value for p3.person.first is an empty string
}
