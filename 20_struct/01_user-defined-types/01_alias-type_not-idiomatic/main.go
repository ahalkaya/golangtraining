package main

import "fmt"

type foo int

func main() {

	var myAge foo
	fmt.Println(myAge) // 0
	myAge = 30
	fmt.Printf("%T %v \n", myAge, myAge) // main.foo 30

}
