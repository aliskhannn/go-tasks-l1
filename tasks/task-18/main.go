package main

import (
	"fmt"
	"sync"
)

type counter struct {
	mu    sync.Mutex
	count int
}

func (c *counter) inc(wg *sync.WaitGroup) {
	defer wg.Done()

	c.mu.Lock()
	c.count++
	fmt.Println("Incrementing", c.count)
	c.mu.Unlock()
}

func (c *counter) printCount() {
	fmt.Printf("Count: %d\n", c.count)
}

func main() {
	var wg sync.WaitGroup

	c := counter{count: 0}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go c.inc(&wg)
	}

	wg.Wait() // wait for all goroutines to finish

	c.printCount()
}
