package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Остановка горутины.
// Реализовать все возможные способы остановки выполнения горутины.
//
// Классические подходы: выход по условию, через канал уведомления,
// через контекст, прекращение работы runtime.Goexit() и др.
//
// Продемонстрируйте каждый способ в отдельном фрагменте кода.

// This program demonstrates various ways to stop a goroutine in Go.
// It includes examples of stopping a goroutine by checking a condition,
// using a channel, using context with cancellation, timeout, and deadline,
// using channel closure, and using runtime.Goexit().

// 1. Example of stopping a goroutine by checking a condition
func stopGoroutineByCondition(wg *sync.WaitGroup) {
	done := false // Shared variable to signal the goroutine to stop

	wg.Add(1) // Increment the WaitGroup counter
	go func() {
		defer wg.Done() // Decrement the WaitGroup counter when the goroutine finishes

		for {
			// Check the condition to stop the goroutine
			if done {
				fmt.Print("Condition-based goroutine stopped\n\n")
				return // Exit the goroutine
			}

			fmt.Println("Condition-based goroutine is running")
			time.Sleep(1 * time.Second) // Simulate work
		}
	}()

	time.Sleep(3 * time.Second) // Let the goroutine run for a while
	done = true                 // Signal the goroutine to stop
	wg.Wait()                   // Wait for the goroutine to finish before returning (for demonstration purposes)
}

// 2. Example of stopping a goroutine using a channel
func stopGoroutineByChannel(wg *sync.WaitGroup) {
	stopChan := make(chan bool, 1) // Channel to signal the goroutine to stop

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-stopChan:
				fmt.Print("Channel-based goroutine stopped\n\n")
				return // Exit the goroutine
			default:
				fmt.Println("Channel-based goroutine is running")
				time.Sleep(1 * time.Second) // Simulate work
			}
		}
	}()

	time.Sleep(3 * time.Second) // Let the goroutine run for a while
	stopChan <- true            // Signal the goroutine to stop
	wg.Wait()                   // Wait for the goroutine to finish before returning (for demonstration purposes)
}

// 3. Example of stopping a goroutine using context with cancellation
func stopGoroutineByContextWithCancel(wg *sync.WaitGroup) {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Print("Context with cancellation-based goroutine stopped\n\n")
				return // Exit the goroutine
			default:
				fmt.Println("Context with cancellation-based goroutine is running")
				time.Sleep(1 * time.Second) // Simulate work
			}
		}
	}()

	time.Sleep(3 * time.Second) // Let the goroutine run for a while
	cancel()                    // Signal the goroutine to stop
	wg.Wait()                   // Wait for the goroutine to finish before returning (for demonstration purposes)
}

// 4. Example of stopping a goroutine using context with timeout
func stopGoroutineByContextWithTimeout(wg *sync.WaitGroup) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Ensure the context is cancelled to free resources

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Print("Context with timeout-based goroutine stopped\n\n")
				return // Exit the goroutine
			default:
				fmt.Println("Context with timeout-based goroutine is running")
				time.Sleep(1 * time.Second) // Simulate work
			}
		}
	}()

	wg.Wait() // Wait for the goroutine to finish before returning (for demonstration purposes)
}

// 5. Example of stopping a goroutine using context with deadline
func stopGoroutineByContextWithDeadline(wg *sync.WaitGroup) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancel() // Ensure the context is cancelled to free resources

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Print("Context with deadline-based goroutine stopped\n\n")
				return // Exit the goroutine
			default:
				fmt.Println("Context with deadline-based goroutine is running")
				time.Sleep(1 * time.Second) // Simulate work
			}
		}
	}()

	wg.Wait() // Wait for the goroutine to finish before returning (for demonstration purposes)
}

// 6. Example of stopping a goroutine using channel closure
func stopGoroutineByChannelClosure(wg *sync.WaitGroup) {
	stopChan := make(chan struct{}) // Channel to signal the goroutine to stop

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-stopChan:
				fmt.Print("Channel closure-based goroutine stopped\n\n")
				return // Exit the goroutine
			default:
				fmt.Println("Channel closure-based goroutine is running")
				time.Sleep(1 * time.Second) // Simulate work
			}
		}
	}()

	time.Sleep(3 * time.Second) // Let the goroutine run for a while
	close(stopChan)             // Signal the goroutine to stop by closing the channel
	wg.Wait()                   // Wait for the goroutine to finish before returning (for demonstration purposes)
}

// 7. Example of stopping a goroutine using runtime.Goexit()
func stopGoroutineByGoexit(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Print("Goexit-based goroutine stopped\n\n")
			wg.Done()
		}()

		for {
			fmt.Println("Goexit-based goroutine is running")
			time.Sleep(3 * time.Second) // Simulate work

			runtime.Goexit() // Stop the goroutine immediately
		}
	}()

	wg.Wait() // Wait for the goroutine to finish before returning (for demonstration purposes)
}

func main() {
	var wg sync.WaitGroup

	// Example 1: Stop goroutine by condition
	stopGoroutineByCondition(&wg)

	// Example 2: Stop goroutine by channel
	stopGoroutineByChannel(&wg)

	// Example 3: Stop goroutine by context with cancellation
	stopGoroutineByContextWithCancel(&wg)

	// Example 4: Stop goroutine by context with timeout
	stopGoroutineByContextWithTimeout(&wg)

	// Example 5: Stop goroutine by context with deadline
	stopGoroutineByContextWithDeadline(&wg)

	// Example 6: Stop goroutine by channel closure
	stopGoroutineByChannelClosure(&wg)

	// Example 7: Stop goroutine by runtime.Goexit()
	stopGoroutineByGoexit(&wg)

	fmt.Println("Main goroutine finished")
}
