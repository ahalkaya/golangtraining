package main

import "fmt"

func main() {

	greetings := map[string]string{
		"tr": "Günaydın",
		"en": "Good morning",
		"fr": "Bonjour",
		"es": "Buenos dias",
		"it": "Bongiorno",
	}

	// Notice that maps are unordered
	fmt.Println(greetings)  // map[es:Buenos dias it:Bongiorno tr:Günaydın en:Good morning fr:Bonjour]
	delete(greetings, "fr") // map[en:Good morning es:Buenos dias it:Bongiorno tr:Günaydın]
	fmt.Println(greetings)

}
