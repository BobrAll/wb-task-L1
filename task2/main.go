package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	PrintSquaresUnordered(nums)
	fmt.Println("\n---")
	PrintSquaresOrdered(nums)
}

// PrintSquaresUnordered prints nums from array without order
func PrintSquaresUnordered(nums []int) {
	var wg sync.WaitGroup
	wg.Add(len(nums))

	for n := range nums {
		go func(n int) {
			defer wg.Done()
			fmt.Printf("%d ", n*n)
		}(n)
	}

	wg.Wait()
}

// PrintSquaresOrdered prints nums from array ordered. Works slower than PrintSquaresUnordered because required synchronization
func PrintSquaresOrdered(nums []int) {
	squares := make([]int, len(nums))
	var wg sync.WaitGroup
	wg.Add(len(nums))

	for i, n := range nums {
		go func(i int, n int) {
			defer wg.Done()
			squares[i] = n * n
		}(i, n)
	}

	wg.Wait()
	for n := range squares {
		fmt.Printf("%d ", n)
	}
}
