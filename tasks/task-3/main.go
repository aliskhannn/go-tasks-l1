package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(id int, dataCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range dataCh {
		fmt.Println("Worker", id, "received data:", data)
		time.Sleep(1 * time.Second) // Simulate some work
	}
	fmt.Println("Worker", id, "exiting gracefully")
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the number of workers. Example: go run main.go 3")
	}

	var wg sync.WaitGroup

	dataCh := make(chan int)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	numOfWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numOfWorkers <= 0 {
		log.Fatal("Please provide a valid positive number of workers")
	}

	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go worker(i+1, dataCh, &wg)
	}

	i := 0
	for {
		select {
		case <-stop:
			// Handle graceful shutdown
			fmt.Println("\nReceived shutdown signal, exiting gracefully...")
			close(dataCh)
			wg.Wait() // Wait for all workers to finish
			return
		default:
			i++
			dataCh <- i * i                    // Send square of i to the channel
			time.Sleep(500 * time.Millisecond) // Simulate work
		}
	}
}
