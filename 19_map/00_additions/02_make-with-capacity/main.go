package main

import "fmt"

func main() {

	user := make(map[string]string, 2)

	fmt.Println(user)      // map[]
	fmt.Println(len(user)) // 0

	user["username"] = "ahalkaya"
	user["password"] = "secret"

	fmt.Println(user)      // map[username:ahalkaya password:secret]
	fmt.Println(len(user)) // 2

	// Notice that length is dynamically incresed
	user["email"] = "example@example.com"
	fmt.Println(user)      // map[username:ahalkaya password:secret email:example@example.com]
	fmt.Println(len(user)) // 3
}
