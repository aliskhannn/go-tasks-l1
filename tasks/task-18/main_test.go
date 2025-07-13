package main

import (
	"sync"
	"testing"
)

func TestCounterConcurrentIncrement(t *testing.T) {
	const goroutines = 1000
	var wg sync.WaitGroup
	c := counter{count: 0}

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go c.inc(&wg)
	}

	wg.Wait()

	if c.count != goroutines {
		t.Errorf("expected count to be %d, got %d", goroutines, c.count)
	}
}
