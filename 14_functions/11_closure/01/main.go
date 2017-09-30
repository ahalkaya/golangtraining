// Define a new scope inside the main function
// and show that a variable outside of that scope
// can be reached while reverse cannot
package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "The credit belongs with the one who is in the ring."
		fmt.Println(y)
	}
	fmt.Print(y) // outside scope of y
}
