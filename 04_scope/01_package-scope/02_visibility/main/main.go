package main

import (
	"fmt"

	"github.com/ahalkaya/golangtraining/04_scope/01_package-scope/02_visibility/vis"
) // imports are file level scope

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
