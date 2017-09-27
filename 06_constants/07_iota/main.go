package main

import "fmt"

// << - bitwise operator to use shifting bits

// 1kb = 1024 bytes = 2^10               bytes
// 1mb = 1024 kb    = 2^10 * 2^10        bytes = 2^20 bytes
// 1gb = 1024 mb    = 2^10 * 2^10 * 2^10 bytes = 2^30 bytes
// 1tb = 1024 gb    = 2^10 * 2^10 * 2^10 bytes = 2^40 bytes

const (
	_  = iota             // 0
	kb = 1 << (iota * 10) // 1 << (1 * 10)
	mb = 1 << (iota * 10) // 1 << (2 * 10)
	gb = 1 << (iota * 10) // 1 << (3 * 10)
	tb = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", kb)
	fmt.Printf("%d\n", kb)
	fmt.Printf("%b\t", mb)
	fmt.Printf("%d\n", mb)
	fmt.Printf("%b\t", gb)
	fmt.Printf("%d\n", gb)
	fmt.Printf("%b\t", tb)
	fmt.Printf("%d\n", tb)
}
