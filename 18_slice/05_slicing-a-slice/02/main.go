package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	fmt.Print("[1:2] ")
	fmt.Println(greeting[1:2]) // [Bonjour!]
	fmt.Print("[:2] ")
	fmt.Println(greeting[:2]) // [Good morning! Bonjour!]
	fmt.Print("[5:] ")
	fmt.Println(greeting[5:]) // [Selamat pagi! Gutten morgen!]
	fmt.Print("[:] ")         // [Good morning! Bonjour! dias! Bongiorno! Ohayo! Selamat pagi! Gutten morgen!]
	fmt.Println(greeting[:])
}
