package main

import (
	"testing"
)

func TestDeleteElement(t *testing.T) {
	tests := []struct {
		s        []int
		i        int
		expected []int
	}{
		{
			s:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			i:        3,
			expected: []int{1, 2, 3, 5, 6, 7, 8, 9, 10},
		},
		{
			s:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			i:        6,
			expected: []int{1, 2, 3, 4, 5, 6, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		result := DeleteElement(tt.s, tt.i)
		for i := range result {
			if result[i] != tt.expected[i] {
				t.Errorf("DeleteElement(%d, %d) = %d; want %d", tt.s, tt.i, result[i], i)
			}
		}
	}
}
