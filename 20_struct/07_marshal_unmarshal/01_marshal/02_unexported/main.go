package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{"James", "Bond", 29}
	bs, _ := json.Marshal(p1)
	fmt.Println(p1)         // {James Bond 29}
	fmt.Println(string(bs)) // {}

}
