/*
For every character in ASCII table print the remainder when divided by 12.
We use this later for our hashing algorithm.
*/
package main

import "fmt"

func main() {

	for i := 65; i <= 122; i++ {
		fmt.Println(i, " - ", string(i), " - ", i%12)
	}

}
