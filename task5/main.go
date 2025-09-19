package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const DefaultWorkingSeconds = 3

// main sends numbers to the channel until the timer expires.
func main() {
	numCh := make(chan int)
	stopTimer := time.After(getWorkingTime())
	go printNums(numCh)

	for i := 0; ; i++ {
		select {
		case <-stopTimer:
			close(numCh)
			return
		case numCh <- i:
		}
	}
}

// getWorkingTime returns the duration from command line or default.
func getWorkingTime() time.Duration {
	var seconds int
	if len(os.Args) < 2 {
		log.Printf("Program was started without specified seconds of working time. Set default value - %d sec."+
			"\nIf you want to specify - run using 'go run task5/main.go {seconds}'", DefaultWorkingSeconds)
		seconds = DefaultWorkingSeconds
	} else {
		seconds = parseSeconds(os.Args[1])
	}
	return time.Duration(seconds) * time.Second
}

// parseSeconds returns seconds from console argument as num
func parseSeconds(s string) int {
	seconds, err := strconv.Atoi(s)
	if err != nil || seconds < 0 {
		log.Printf("Invalid input, seconds must be a positive num. Set default value - %d sec.\n", DefaultWorkingSeconds)
		seconds = DefaultWorkingSeconds
	}
	return seconds
}

// printNums reads numbers from the channel and prints them until the channel is closed.
func printNums(numCh <-chan int) {
	for num := range numCh {
		fmt.Println(num)
	}
}
