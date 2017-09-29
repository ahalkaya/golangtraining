package main

import "fmt"

func main() {

	switch "Go" {
	case "Java":
		fmt.Println("Java is cool")
	case "PHP":
		fmt.Println("PHP is cool")
	case "Go":
		fmt.Println("Go is cool")
		fallthrough
	case "JavaScript":
		fmt.Println("JavaScript is also cool")
		fallthrough
	case "Phython":
		fmt.Println("Phython is also cool")
	case "Ruby":
		fmt.Println("Ruby is cool")
	}

}

/*
  no default fallthrough
  fallthrough is optional
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/
