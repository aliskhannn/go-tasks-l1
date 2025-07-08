package main

import (
	"fmt"
	"sync"
)

// generate sends numbers from the slice to the `out` channel, then closes the channel
func generate(nums []int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Marks this goroutine as done when the function returns

	for _, v := range nums {
		out <- v // Send each number to the output channel
	}
	close(out) // Close the channel to signal that no more values will be sent
}

// multiplyByTwo reads numbers from `in`, doubles them, and sends them to `out`
func multiplyByTwo(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range in { // Reading until the input channel is closed
		out <- v * 2 // Multiply and send the result
	}
	close(out) // Close the output channel when done
}

// printResults reads from the channel and prints the values
func printResults(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range in {
		fmt.Println(v) // Print each value received
	}
}

func main() {
	var wg sync.WaitGroup

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch1 := make(chan int, len(nums))
	ch2 := make(chan int, len(nums))

	wg.Add(3)

	// Launch the pipeline stages
	go generate(nums, ch1, &wg)     // Stage 1: generate numbers
	go multiplyByTwo(ch1, ch2, &wg) // Stage 2: multiply by 2
	go printResults(ch2, &wg)       // Stage 3: print the result

	wg.Wait() // Wait for all goroutines to finish
}
