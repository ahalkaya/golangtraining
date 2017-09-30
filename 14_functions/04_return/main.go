// Declare a function which returns a string
package main

import "fmt"

func main() {
	fmt.Println(greetFullName("Jane ", "Doe"))
	fmt.Println(greetFullNameF("Jane", "Doe"))
}

func greetFullName(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

func greetFullNameF(fname, lname string) string {
	return fname + " " + lname
}
