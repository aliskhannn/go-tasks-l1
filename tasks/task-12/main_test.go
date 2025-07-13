package main

import (
	"reflect"
	"testing"
)

func TestStringSet(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"cat", "cat", "dog", "dog", "tree"}, []string{"cat", "dog", "tree"}},
		{[]string{}, []string{}},
		{[]string{"a", "b", "a", "c", "b"}, []string{"a", "b", "c"}},
		{[]string{"unique"}, []string{"unique"}},
	}

	for _, tt := range tests {
		result := StringSet(tt.input)

		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("StringSet(%v) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}
