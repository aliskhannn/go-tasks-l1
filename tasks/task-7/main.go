package main

import (
	"fmt"
	"sync"
)

// Конкурентная запись в map.
// Реализовать безопасную для конкуренции запись данных в структуру map.
//
// Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).
//
// Проверьте работу кода на гонки (util go run -race).

// 1. Example of using sync.Map for concurrent writes
func exampleSyncMap(wg *sync.WaitGroup) {
	fmt.Println("exampleSyncMap called")

	var m sync.Map

	for i := 0; i < 100; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go func() {
			defer wg.Done() // Decrement the counter when the goroutine completes
			m.Store(i, i*i) // Store a value in the map
		}()
	}

	wg.Wait() // Wait for all goroutines to finish

	for i := 0; i < 100; i++ {
		value, ok := m.Load(i) // Load a value from the map
		if ok {
			fmt.Println("Key:", i, "Value:", value.(int)) // Print the key and value
		} else {
			fmt.Println("Key:", i, "not found")
		}
	}

	fmt.Println("exampleSyncMap completed")
}

// 2. Example of using built-in map with sync.Mutex for concurrent writes
func exampleMutexMap(wg *sync.WaitGroup, mu *sync.Mutex) {
	fmt.Println("exampleMutexMap called")

	var m = make(map[int]int)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i * i // Safely write to the map
			mu.Unlock()
		}()
	}

	wg.Wait() // Wait for all goroutines to finish

	for i := 0; i < 100; i++ {
		mu.Lock()
		value, ok := m[i] // Safely read from the map
		mu.Unlock()
		if ok {
			fmt.Println("Key:", i, "Value:", value) // Print the key and value
		} else {
			fmt.Println("Key:", i, "not found")
		}
	}

	fmt.Println("exampleMutexMap completed")
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Example of using sync.Map for concurrent writes
	// Uncomment the line below to run the sync.Map example
	// exampleSyncMap(&wg)

	// Example of using built-in map with sync.Mutex for concurrent writes
	exampleMutexMap(&wg, &mu)
}
