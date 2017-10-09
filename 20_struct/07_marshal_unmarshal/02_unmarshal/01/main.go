package main

import (
	"encoding/json"
	"fmt"
)

/*
JSON Unmarshal
In this example we convert a byte slice into a struct.
*/

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	var p1 person
	fmt.Println(p1.First) //
	fmt.Println(p1.Last)  //
	fmt.Println(p1.Age)   // 0

	// convertion of a string into a byte slice
	bs := []byte(`{"First":"James", "Last":"Bond", "Age": 29}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("--------------------")
	fmt.Println(p1.First)   // James
	fmt.Println(p1.Last)    // Bond
	fmt.Println(p1.Age)     // 29
	fmt.Printf("%T \n", p1) // main.person

}
