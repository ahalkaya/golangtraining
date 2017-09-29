package main

import "fmt"

func main() {

	planet := "Mars"

	switch {
	case len(planet) == 3:
		fmt.Println("The planet is Sun")
	case planet == "Mars":
		fmt.Println("The planet is Mars")
	case planet == "Jupiter":
		fmt.Println("The planet is Jupiter")
	case planet == "Mars", planet == "Saturn", planet == "Uranus":
		fmt.Println("The planet is either Saturn or Uranus")
	case planet == "Neptune":
		fmt.Println("The planet is Neptune")
	case planet == "Venus":
		fmt.Println("The planet is Venus")
	default:
		fmt.Println("The default planet is Earth")
	}

}

/*
  expression not needed
  -- if no expression provided, go checks for the first case that evals to true
  -- makes the switch operate like if/if else/else
  cases can be expressions
*/
