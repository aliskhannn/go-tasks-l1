package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		s        []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8, 7},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{}, 1, -1},
	}

	for _, tt := range tests {
		got := binarySearch(tt.s, tt.target)
		if got != tt.expected {
			t.Errorf("binarySearch(%v, %d) = %d; want %d", tt.s, tt.target, got, tt.expected)
		}
	}
}
