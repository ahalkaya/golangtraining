package main

import "fmt"

func main() {

	letterA := "A"[0]
	fmt.Println(letterA)         // 65
	fmt.Printf("%T \n", letterA) // uint8

	letterB := rune("B"[0])
	fmt.Println(letterB)         // 66
	fmt.Printf("%T \n", letterB) // int32
}
