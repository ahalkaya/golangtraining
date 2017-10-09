package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	Id       int
	Name     string
	Username string
	Email    string
	Address  address
	Phone    string
	Website  string
	Company  company
}

type address struct {
	Street  string
	Suite   string
	City    string
	Zipcode string
	Geo     geo
}

type geo struct {
	Lat string
	Lng string
}

type company struct {
	Name        string
	CatchPhrase string
	Bs          string
}

func main() {

	var u1 []user
	res, err := http.Get("http://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	errr := json.NewDecoder(res.Body).Decode(&u1)
	if errr != nil {
		log.Fatal(errr)
	}
	fmt.Println(u1[0])

}

/*
http://jsonplaceholder.typicode.com/users

{
	id: 1,
	name: "Leanne Graham",
	username: "Bret",
	email: "Sincere@april.biz",
	address: {
		street: "Kulas Light",
		suite: "Apt. 556",
		city: "Gwenborough",
		zipcode: "92998-3874",
		geo: {
			lat: "-37.3159",
			lng: "81.1496"
		}
	},
	phone: "1-770-736-8031 x56442",
	website: "hildegard.org",
	company: {
		name: "Romaguera-Crona",
		catchPhrase: "Multi-layered client-server neural-net",
		bs: "harness real-time e-markets"
	}
}

*/
