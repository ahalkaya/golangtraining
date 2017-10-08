/*
A rune is an alias to int32.
A rune is a slice of bytes.
A rune is a number.
*/

package main

import "fmt"

func main() {

	letter := 'A'
	fmt.Println(letter)         // 65
	fmt.Printf("%T \n", letter) // int32

}
