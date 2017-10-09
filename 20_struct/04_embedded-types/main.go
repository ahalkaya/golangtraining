package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type doubleZero struct {
	person
	licenseToKill bool
}

func main() {

	p1 := doubleZero{
		person: person{
			first: "James",
			last:  "Bond",
			age:   29,
		},
		licenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			first: "Miss",
			last:  "MoneyPenny",
			age:   25,
		},
		licenseToKill: false,
	}

	fmt.Println(p1.first, p1.last, p1.licenseToKill)
	fmt.Println(p2.first, p2.last, p2.licenseToKill)
}
