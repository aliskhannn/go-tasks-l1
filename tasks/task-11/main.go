package main

import "fmt"

func getIntersection(a, b []int) []int {
	m := make(map[int]bool)    // to save elems from 'a'
	seen := make(map[int]bool) // to avoid it duplicated
	result := make([]int, 0)

	// fill the map with elements from 'a'
	for _, v := range a {
		m[v] = true
	}

	for _, v := range b {
		if m[v] && !seen[v] { // If an element exists in 'a' and not already added
			seen[v] = true // mark as added
			result = append(result, v)
		}
	}

	return result
}

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	intersection := getIntersection(a, b)
	fmt.Println(intersection)
}
