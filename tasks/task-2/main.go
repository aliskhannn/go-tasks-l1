package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	nums := [5]int{2, 4, 6, 8, 10}
	result := make([]int, len(nums))

	for i, num := range nums {
		wg.Add(1)
		go func(i, num int) {
			defer wg.Done()
			result[i] = num * num
		}(i, num)
	}

	wg.Wait()

	fmt.Println("Results:", result)
}
