package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	var p1 person
	reader := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}`)
	// Decode reads the next JSON-encoded value from its
	// input and stores it in the value pointed to by v.
	json.NewDecoder(reader).Decode(&p1)

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

}
