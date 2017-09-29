/*
Create a program that prints to the terminal asking for
a user to enter a small number and a larger number.
Print the remainder of the larger number
divided by the smaller number.
*/

package main

import "fmt"

func main() {

	var numOne int
	var numTwo int
	fmt.Print("Enter a large number: ")
	fmt.Scan(&numOne)
	fmt.Print("Enter a smaller number: ")
	fmt.Scan(&numTwo)
	fmt.Printf("%d %% %d = %d \n", numOne, numTwo, numOne%numTwo)
}
