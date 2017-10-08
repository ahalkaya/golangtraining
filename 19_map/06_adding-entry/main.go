package main

import "fmt"

func main() {

	greetings := map[string]string{
		"tr": "Günaydın",
		"en": "Good morning",
		"fr": "Bonjour",
	}

	greetings["es"] = "Buenos días"

	fmt.Println(greetings) // map[tr:Günaydın en:Good morning fr:Bonjour es:Buenos días]

}
