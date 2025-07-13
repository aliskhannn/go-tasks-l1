package main

import "testing"

func TestReverseSting(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"главрыба", "абырвалг"},
		{"главрыба😄", "😄абырвалг"},
		{"", ""},
		{"а", "а"},
		{"Go🚀Lang", "gnaL🚀oG"},
		{"12345", "54321"},
	}

	for _, tt := range tests {
		result := reverseSting(tt.input)
		if result != tt.expected {
			t.Errorf("reverseSting(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
