package main

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"snow dog sun", "sun dog snow"},
		{"hello world", "world hello"},
		{"one", "one"},
		{"", ""},
		{"a b c", "c b a"},
		{"  leading and trailing  ", "  trailing and leading  "}, // пробелы остаются
		{"go is fun", "fun is go"},
	}

	for _, tt := range tests {
		result := reverseWords(tt.input)
		if result != tt.expected {
			t.Errorf("reverseWords(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
