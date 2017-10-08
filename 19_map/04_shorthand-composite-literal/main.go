package main

import "fmt"

func main() {

	greetings := map[string]string{}

	fmt.Println(greetings)        // map[]
	fmt.Println(greetings == nil) // false

	// Since the map is not nil, below code works
	greetings["en"] = "Good morning"
	greetings["fr"] = "Bonjour"

	fmt.Println(greetings) // map[en:Good morning fr:Bonjour]

}
