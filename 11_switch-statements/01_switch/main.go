package main

import "fmt"

func main() {

	switch "blue" {
	case "green":
		fmt.Println("My favorite color is green")
	case "red":
		fmt.Println("My favorite color is red")
	case "blue":
		fmt.Println("My favorite color is blue")
	default:
		fmt.Println("What is your favorite color?")
	}

}

/*
  no default fallthrough
  fallthrough is optional
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/
