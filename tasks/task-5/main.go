package main

import (
	"fmt"
	"time"
)

// Таймаут на канал.

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.

func main() {
	ch := make(chan int)
	done := make(chan struct{}) // channel for signal goroutine to finish
	// Timeout channel to stop the program after 5 seconds
	timeout := time.After(5 * time.Second)

	go func() {
		i := 1
		for {
			select {
			case <-done:
				fmt.Println("Stopping the goroutine.")
				return
			case ch <- i * i: // Send square of i to the channel
				i++
				time.Sleep(1 * time.Second) // Simulate some work
			}
		}
	}()

	for {
		select {
		case val := <-ch:
			fmt.Println("Received:", val)
		case <-timeout:
			fmt.Println("Timeout reached.")
			fmt.Println("Stopping the program.")
			close(done)
			time.Sleep(100 * time.Millisecond) // Wait for goroutine to finish
			return
		}
	}
}
