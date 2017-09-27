package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// blank identifier is mostly used in return type situations
	// here it denotes that returning error is not used
	res, _ := http.Get("http://jsonplaceholder.typicode.com/users")
	// returning error is not used
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
