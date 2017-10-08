package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting) // map[0:Good morning! 1:Bonjour! 2:Buenos dias! 3:Bongiorno!]

	delete(myGreeting, 7) // There is no entry with key 7 but that gives no error

	fmt.Println(myGreeting) // map[0:Good morning! 1:Bonjour! 2:Buenos dias! 3:Bongiorno!]
}
