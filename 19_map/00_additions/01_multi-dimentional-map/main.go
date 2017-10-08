package main

import "fmt"

func main() {

	users := map[string]map[string]string{
		"jondoe": map[string]string{
			"fname":   "Jon",
			"lname":   "Doe",
			"country": "us",
		},
		"janedoe": map[string]string{
			"fname":   "Jane",
			"lname":   "Doe",
			"country": "gb",
		},
	}

	fmt.Println(users)

	for k, v := range users {
		fmt.Println("User with nickname:", k)
		fmt.Println("First name:", v["fname"])
		fmt.Println("Last name:", v["lname"])
		fmt.Println("Country:", v["country"])
	}

}
