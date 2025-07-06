package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// Завершение по Ctrl+C
// Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
//
// Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.

func worker(ctx context.Context, id int, dataCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		// Check if the context is done (e.g., due to cancellation)
		case <-ctx.Done():
			fmt.Println("Worker", id, "exiting gracefully due to context cancellation")
			return // Exit the worker if the context is done
		// Read data from the channel
		case data, ok := <-dataCh:
			// If the channel is closed, exit the worker
			if !ok {
				fmt.Println("Worker", id, "exiting gracefully due to channel closure")
				return
			}
			fmt.Println("Worker", id, "received data:", data)
			time.Sleep(1 * time.Second) // Simulate some work
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the number of workers. Example: go run main.go 3")
	}

	var wg sync.WaitGroup

	dataCh := make(chan int)

	// To handle signals, here I use a context, which is the preferred approach for graceful shutdown in Go.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop() // Ensure the context is cancelled on exit

	numOfWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numOfWorkers <= 0 {
		log.Fatal("Please provide a valid positive number of workers")
	}

	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i+1, dataCh, &wg)
	}

	i := 0
	for {
		select {
		case <-ctx.Done():
			// Handle graceful shutdown
			fmt.Println("\nReceived shutdown signal, exiting gracefully...")
			close(dataCh) // Close the channel to signal workers to exit
			wg.Wait()     // Wait for all workers to finish
			return
		default:
			i++
			dataCh <- i * i                    // Send square of i to the channel
			time.Sleep(500 * time.Millisecond) // Simulate work
		}
	}
}
