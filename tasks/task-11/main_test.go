package main

import (
	"reflect"
	"testing"
)

func TestGetIntersection(t *testing.T) {
	tests := []struct {
		a, b     []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{2, 3, 4}, []int{2, 3}},
		{[]int{1, 2, 2, 3}, []int{2, 2, 3, 5}, []int{2, 3}},
		{[]int{1, 4, 5}, []int{2, 3, 6}, []int{}},
		{[]int{}, []int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3}, []int{}, []int{}},
		{[]int{}, []int{}, []int{}},
		{[]int{1, 1, 1}, []int{1, 1}, []int{1}},
	}

	for _, tt := range tests {
		result := getIntersection(tt.a, tt.b)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("getIntersection(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
		}
	}
}
