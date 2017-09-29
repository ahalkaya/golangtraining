package main

import "fmt"

func main() {

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

}

/*
// Keeping scope tight - scope of initialized variable is if block
// initialization         ; expression
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
*/
