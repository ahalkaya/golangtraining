package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Create slice to hold counts
	buckets := make([]int, 200)
	// Loop over the words
	for scanner.Scan() {
		n := hashBucket(scanner.Text())
		buckets[n]++
	}
	fmt.Println("***********")
	for i := 0; i < len(buckets); i++ {
		if token := string(i); buckets[i] > 0 {
			fmt.Printf("i: %d - buckets[i]: %d - token: %s\n", i, buckets[i], token)
		}
	}
}

func hashBucket(word string) int {
	return int(word[0])
}
