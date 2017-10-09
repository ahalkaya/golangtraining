package main

import "fmt"

type foo int

func main() {

	var myAge foo
	myAge = 30
	fmt.Printf("%T %v \n", myAge, myAge) // main.foo 30

	var yourAge int
	yourAge = 25
	fmt.Printf("%T %v \n", yourAge, yourAge) // int 25

	// this doesn't work:
	//	 fmt.Println(myAge + yourAge)

	// this conversion works:
	//	 fmt.Println(int(myAge) + yourAge)

}
