package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func captureOutput(f func()) string {
	old := os.Stdout // сохраняем оригинал
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	os.Stdout = old // восстанавливаем оригинал

	return buf.String()
}

func TestIdentifyType(t *testing.T) {
	cInt := make(chan int)
	cBool := make(chan bool)

	tests := []struct {
		input    interface{}
		expected string
	}{
		{1, "int: 1\n"},
		{"hello", "string: hello\n"},
		{true, "bool: true\n"},
		{cInt, fmt.Sprintf("chan: %v\n", reflect.TypeOf(cInt))},
		{cBool, fmt.Sprintf("chan: %v\n", reflect.TypeOf(cBool))},
		{12.3, "unknown type: 12.3\n"},
	}

	for _, tt := range tests {
		out := captureOutput(func() {
			identifyType(tt.input)
		})

		if out != tt.expected {
			t.Errorf("identifyType(%v) output = %q, want %q", tt.input, out, tt.expected)
		}
	}
}
