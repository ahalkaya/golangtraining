package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)
	buckets := make([][]string, 12)
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}

	// Print len of each bucket
	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}
	// Print the words in one of the buckets
	// fmt.Println(buckets[6])
	fmt.Println(len(buckets))
	fmt.Println(cap(buckets))

}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
	// a more uneven distribution
	// return len(word) % buckets
}
