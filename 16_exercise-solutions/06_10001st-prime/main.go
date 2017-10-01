/*
Project Euler - Problem 7
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
we can see that the 6th prime is 13.
What is the 10 001st prime number?
*/
package main

import "fmt"

func main() {

	var index int
	prime := 1
	for index < 10001 {
		prime++
		if isPrime(prime) {
			index++
		}
	}
	fmt.Println(prime) // 104743

}

func isPrime(number int) bool {
	var count int
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			count++
			break
		}
	}
	return count == 0
}
