package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 3, 6, 8, 4, 9, 2, 7, 1, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{3, 3, 3}, []int{3, 3, 3}},
		{[]int{10, 9, 8, 7, 6}, []int{6, 7, 8, 9, 10}},
	}

	for _, tt := range tests {
		got := quickSort(tt.input)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("quickSort(%v) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}
