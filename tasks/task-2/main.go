package main

import (
	"fmt"
	"sync"
)

// Конкурентное возведение в квадрат.
//
// Написать программу, которая конкурентно рассчитает значения квадратов чисел,
// взятых из массива [2,4,6,8,10], и выведет результаты в stdout.

func main() {
	var wg sync.WaitGroup

	nums := [5]int{2, 4, 6, 8, 10}
	result := make([]int, len(nums))

	for i, num := range nums {
		// Increment the WaitGroup counter for each goroutine
		wg.Add(1)
		go func() {
			defer wg.Done()
			result[i] = num * num
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Results:", result)
}
