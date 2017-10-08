package main

import "fmt"

func main() {

	greetings := make(map[string]string)

	fmt.Println(greetings)        // map[]
	fmt.Println(greetings == nil) // false

	greetings["en"] = "Good morning"
	greetings["fr"] = "Bonjour"

	fmt.Println(greetings) // map[en:Good morning fr:Bonjour]

}
