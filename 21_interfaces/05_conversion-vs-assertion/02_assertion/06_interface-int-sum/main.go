// https://golang.org/ref/spec#Type_assertions
package main

import "fmt"

func main() {

	var a interface{} = 7 // a has dynamic type int and value 7
	var b = 5

	sum := a.(int) + b

	fmt.Println(sum) // 12

}
