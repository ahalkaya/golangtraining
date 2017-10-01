/*
Modify the previous program (01_half) to use a func expression.
*/
package main

import "fmt"

func main() {
	half := func(i int) (float64, bool) {
		return float64(i) / 2, i%2 == 0
	}
	fmt.Println(half(9))
}
