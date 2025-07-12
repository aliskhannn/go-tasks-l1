package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 3

	// Shift the tail to the location of the element being removed
	copy(s[i:], s[i+1:])

	// Zero the last element
	s[len(s)-1] = 0

	// Decrease slice length by 1
	s = s[:len(s)-1]

	fmt.Println(s) // [1 2 3 5 6 7 8 9 10]
}
