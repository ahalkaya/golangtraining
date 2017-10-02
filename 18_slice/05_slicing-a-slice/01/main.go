package main

import "fmt"

func main() {

	mySlice := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Print("[:] ")
	fmt.Println(mySlice[:])
	fmt.Print("[2:4] ")        // index 4 is exluded
	fmt.Println(mySlice[2])    // index access; accessing by index
	fmt.Println("myString"[2]) // index access; accessing by index
	// "myString"[2] prints 83 which is the decimal value
	// corresponding to unicode "S"

}

/*
You can do "slicing" on a string because a string is a slice of bytes.
Remember:
    - A string is made up of runes.
    - A rune is a unicode code point.
    - A unicode code point is 1 to 4 bytes.
Strings are made up of runes, runes are made up of bytes, so strings
are made up of bytes. A string is a bunch of bytes; a slice of bytes.
*/
