// Using an int pointer
package main

import "fmt"

func main() {
	age := 30
	fmt.Println(&age) // 0xc420012088

	changeMe(&age)

	fmt.Println(&age) // 0xc420012088
	fmt.Println(age)  // 24
}

func changeMe(z *int) {
	fmt.Println(z)  // 0xc420012088
	fmt.Println(*z) //30
	*z = 24
	fmt.Println(z)  // 0xc420012088
	fmt.Println(*z) // 24
}
