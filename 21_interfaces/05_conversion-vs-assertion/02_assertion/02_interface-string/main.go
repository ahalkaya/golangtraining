// https://golang.org/ref/spec#Type_assertions
package main

import "fmt"

func main() {
	var name interface{} = "Sydney" // name has dynamic type strint and value Sydney
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str) // string
	} else {
		fmt.Printf("value is not a string\n")
	}
}
