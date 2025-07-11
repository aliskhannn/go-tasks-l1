package main

import "fmt"

// Printer / Target
type Printer interface {
	Print(text string)
}

// OldPrinter / Adaptee
type OldPrinter struct{}

func (op *OldPrinter) PrintOld(text string) {
	fmt.Printf("Old printer: %s\n", text)
}

// PrinterAdapter / Adapter
type PrinterAdapter struct {
	*OldPrinter
}

func (p *PrinterAdapter) Print(text string) {
	p.PrintOld(text)
}

func main() {
	old := &OldPrinter{}
	adapter := &PrinterAdapter{OldPrinter: old}

	var printer Printer = adapter
	printer.Print("Hello") // adapter uses old method PrintOld
}
