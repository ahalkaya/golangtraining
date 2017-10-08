package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words[scanner.Text()] = ""
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0
	for k := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}

}
