package main

import (
	"fmt"
	"time"
)

var workerID int
var publisherID int

func main() {
	input := make(chan string)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	time.Sleep(1 * time.Millisecond)
}

// publisher pushes data into a channel
func publisher(out chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d is pushing data\n", thisID)
		data := fmt.Sprintf("Data from publisher %d. Data %d", thisID, dataID)
		out <- data
	}
}

func workerProcess(in <-chan string) {
	workerID++
	thisID := workerID
	for {
		fmt.Printf("%d: waiting for input...\n", thisID)
		input := <-in
		fmt.Printf("%d: input is: %s\n", thisID, input)
	}
}

/*
CHALLENGE #1
Is this fan out?
CHALLENGE #2
Is this fan in?
*/

/*
FAN OUT: Multiple functions reading from the same channel until that channel is closed.
FAN IN:  A function can read from multiple inputs and proceed until all are closed by
		 multiplexing the input channels onto a single channel that's closed when all
		 the inputs are closed.
*/
