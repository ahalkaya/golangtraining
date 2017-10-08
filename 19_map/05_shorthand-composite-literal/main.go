package main

import "fmt"

func main() {

	languages := []string{
		"en",
		"fr", // Notice the comma at the end
	}

	fmt.Println(languages) // [en fr]

	greetings := map[string]string{
		"en": "Good morning",
		"fr": "Bonjour", // Notice the comma at the end
	}

	fmt.Println(greetings) // map[en:Good morning fr:Bonjour]

}
