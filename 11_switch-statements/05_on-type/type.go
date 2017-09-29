package main

import (
	"fmt"
)

// switch on types
// -- normally we switch on value of variable
// -- go allows you to switch on type of variable

type contact struct {
	greeting string
	name     string
}

// SwitchOnType works with interfaces
// we'll learn more about interfaces later
func switchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	switchOnType(3)
	switchOnType("Hey")
	var t = contact{"Good to see you,", "Alihan"}
	switchOnType(t)
	switchOnType(t.greeting)
	switchOnType(t.name)
}
