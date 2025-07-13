package main

import "testing"

func TestIsUnique(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abcd", true},
		{"abCdefAaf", false},
		{"", true},
		{"AaBbCcDd", false},
		{"abcABC", false},
	}

	for _, tt := range tests {
		result := isUnique(tt.input)
		if result != tt.expected {
			t.Errorf("isUnique(%s) = %t, want %t", tt.input, result, tt.expected)
		}
	}
}
