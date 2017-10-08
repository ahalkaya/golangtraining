package main

import "fmt"

func main() {

	var greetings = make(map[string]string)

	// Capacity is optional
	// var greetings = make(map[string]string, 2)

	fmt.Println(greetings)        // map[]
	fmt.Println(greetings == nil) // false

	greetings["en"] = "Good morning"
	greetings["fr"] = "Bonjour"

	fmt.Println(greetings) // map[en:Good morning fr:Bonjour]

	mySlice := []string{}
	mySlice[0] = "Hey" // runtime error: index out of range
	// Need to use _append_

}
