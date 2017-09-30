package main

import "fmt"

func main() {

	if true && false {
		fmt.Println("This will not run until both conditions are true.")
	}

}
