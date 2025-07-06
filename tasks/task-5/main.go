package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// Timeout channel to stop the program after 5 seconds
	timeout := time.After(5 * time.Second)

	go func() {
		i := 1
		for {
			ch <- i * i // Send square of i to the channel
			i++
			time.Sleep(1 * time.Second) // Simulate some work
		}
	}()

	for {
		select {
		case val := <-ch:
			fmt.Println("Received:", val)
		case <-timeout:
			fmt.Println("Timeout reached, stopping the program.")
			return
		}
	}
}
