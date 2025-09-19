package main

import (
	"fmt"
	"time"
)

func main() {
	numCh := make(chan int)
	timer := time.After(3 * time.Second)
	go printNums(numCh)

	for i := 0; ; i++ {
		select {
		case <-timer:
			close(numCh)
			return
		case numCh <- i:
		}
	}
}

func printNums(numCh <-chan int) {
	for num := range numCh {
		fmt.Println(num)
	}
}
