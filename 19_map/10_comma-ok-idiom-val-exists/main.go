package main

import "fmt"

func main() {

	greetings := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(greetings)

	delete(greetings, 2)

	// comma ok idiom -> expression and condition seperated by comma
	if val, exists := greetings[2]; exists {
		fmt.Println("That value exists.")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	}

}
