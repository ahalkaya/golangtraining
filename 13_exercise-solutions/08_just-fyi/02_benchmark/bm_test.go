package benchmark

import (
	"fmt"
	"testing"
)

func BenchmarkThreeFive(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}

// run this at command:
// go test -bench='.*'
