// https://golang.org/ref/spec#Type_assertions
package main

import "fmt"

func main() {
	var val interface{} = 7
	fmt.Println(val + 6) // invalid operation: val + 6 (mismatched types interface {} and int)
}
