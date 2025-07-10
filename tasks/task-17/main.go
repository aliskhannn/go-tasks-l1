package main

import "fmt"

func binarySearch(s []int, target int) int {
	left, right := 0, len(s)-1

	for left <= right {
		mid := (left + right) / 2

		if s[mid] == target {
			return mid
		}

		if s[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 8

	fmt.Println(binarySearch(s, target))
}
