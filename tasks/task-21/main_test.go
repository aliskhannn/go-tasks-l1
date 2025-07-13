package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestPrinterAdapter(t *testing.T) {
	// Substitute stdout to capture output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Подготовка адаптера
	oldPrinter := &OldPrinter{}
	adapter := &PrinterAdapter{OldPrinter: oldPrinter}
	var printer Printer = adapter

	// Call method
	printer.Print("Hello")

	// Finish recording and read the output
	_ = w.Close()
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	os.Stdout = oldStdout

	output := buf.String()
	expected := "Old printer: Hello\n"

	if output != expected {
		t.Errorf("unexpected output: got %q, want %q", output, expected)
	}
}
