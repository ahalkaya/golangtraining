// https://golang.org/ref/spec#Type_assertions
package main

import "fmt"

func main() {
	var val interface{} = 7
	fmt.Printf("%T\n", val) // int
}
