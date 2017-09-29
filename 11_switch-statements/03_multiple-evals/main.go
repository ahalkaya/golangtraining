package main

import "fmt"

func main() {

	switch "September" {
	case "March", "April", "May":
		fmt.Println("The season is spring")
	case "June", "July", "August":
		fmt.Println("The season is summer")
	case "September", "October", "November":
		fmt.Println("The season is autumn")
	case "December", "January", "February":
		fmt.Println("The season is winter")
	}

}
