package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
}

func main() {

	var p1 person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First":"James", "Last":"Bond", "wisdom score":29}`)
	err := json.Unmarshal(bs, &p1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)

	/*


		0
		--------------
		James
		Bond
		29
		main.person
	*/

}
