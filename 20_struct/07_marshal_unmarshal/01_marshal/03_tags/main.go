package main

import "fmt"
import "encoding/json"

func main() {

	type person struct {
		First string
		Last  string `json:"-"`
		Age   int    `json:"wisdom score"`
	}

	p1 := person{"James", "Bond", 29}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))

}

/*
// Field appears in JSON as key "myName".
Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
Field int `json:",omitempty"`

// Field is ignored by this package.
Field int `json:"-"`

// Field appears in JSON as key "-".
Field int `json:"-,"`
*/
