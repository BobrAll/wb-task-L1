package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

var (
	DefaultWorkersCount = 5
)

func main() {
	workersCount, err := getWorkersCount()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Set default workers count to", DefaultWorkersCount)
		workersCount = DefaultWorkersCount
	}

	numCh := make(chan int)
	stopSignalCh := make(chan struct{})
	wg := &sync.WaitGroup{}

	printInstruction()
	go listenSigint(stopSignalCh)

	runWorkers(wg, workersCount, numCh, stopSignalCh)
	startSpamChannel(numCh, stopSignalCh)
	wg.Wait()
}

// getWorkersCount parses the number of workers from command line arguments.
func getWorkersCount() (int, error) {
	if len(os.Args) < 2 {
		return 0, errors.New("missing argument: count of workers\n" +
			"usage: 'go run task3.main.go {num}'\n" +
			"where num - count of workers")
	}

	workersCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return 0, errors.New("console argument (count of workers) must be a number")
	}

	return workersCount, nil
}

// printInstruction prints usage instruction and waits for initial Enter press.
func printInstruction() {
	fmt.Println("To start printing press 'Enter'. Than press 'Ctrl+C' to stop printing.")
	waitForUserInput()
}

// runWorkers starts N worker goroutines to read from numCh and print values.
func runWorkers(wg *sync.WaitGroup, workersCount int, ch chan int, stopSignalCh chan struct{}) {
	wg.Add(workersCount)
	for i := 0; i < workersCount; i++ {
		go printFromChannel(ch, stopSignalCh, wg)
	}
}

// printFromChannel is a worker goroutine: reads numbers from channel until stop signal or channel closes.
func printFromChannel(numCh chan int, stopSignalCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stopSignalCh:
			return
		case num, ok := <-numCh:
			if !ok {
				return
			}
			fmt.Println(num)
		}
	}
}

// startSpamChannel generates integers endlessly and sends them to the channel until stop signal.
func startSpamChannel(numCh chan int, stopSignalCh chan struct{}) {
	for i := 0; ; i++ {
		select {
		case <-stopSignalCh:
			close(numCh)
			return
		case numCh <- i:
		}
	}
}

// listenSigint waits for sigint to signal all goroutines to stop.
func listenSigint(stopSignalCh chan struct{}) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	close(stopSignalCh)
	signal.Stop(sigCh)
}

// waitForUserInput pauses execution until the user presses Enter.
func waitForUserInput() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _ = reader.ReadLine()
}
