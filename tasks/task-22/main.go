package main

import (
	"fmt"
	"math/big"
	"strings"
)

// add returns the sum of a and b as a new *big.Int.
func add(a, b *big.Int) *big.Int {
	sum := new(big.Int)

	return sum.Add(a, b)
}

// subtract returns the result of subtracting b from a as a new *big.Int.
func subtract(a, b *big.Int) *big.Int {
	res := new(big.Int)

	return res.Sub(a, b)
}

// divide returns the result of dividing a by b as a new *big.Int.
// If b == 0, a division-by-zero run-time panic occurs.
func divide(a, b *big.Int) *big.Int {
	res := new(big.Int)

	return res.Div(a, b)
}

// multiply returns the result of multiplying a and b as a new *big.Int.
func multiply(a, b *big.Int) *big.Int {
	res := new(big.Int)

	return res.Mul(a, b)
}

func main() {
	// Create 2 big integers for math operations
	a := big.NewInt(1_001_234_837_633)
	b := big.NewInt(3_434_392_534)

	// Format numbers and print results
	fmt.Printf("%s + %s = %s\n", formatNumber(a), formatNumber(b), formatNumber(add(a, b)))
	fmt.Printf("%s - %s = %s\n", formatNumber(a), formatNumber(b), formatNumber(subtract(a, b)))
	fmt.Printf("%s / %s = %s\n", formatNumber(a), formatNumber(b), formatNumber(divide(a, b)))
	fmt.Printf("%s * %s = %s\n", formatNumber(a), formatNumber(b), formatNumber(multiply(a, b)))
}

// formatNumber formats a number by adding underscores as a thousand separators and improving readability.
// For example, 3434392 becomes "3_434_392".
func formatNumber(n *big.Int) string {
	s := fmt.Sprintf("%d", n) // convert the number to string
	var b strings.Builder     // efficiently build the formatted string

	// Calculate the number of digits before the first underscore.
	// This helps split the string into groups of three digits from the right.
	remainder := len(s) % 3
	if remainder == 0 {
		remainder = 3
	}

	// Write the first group of digits.
	b.WriteString(s[:remainder])

	// Write the remaining digits in groups of three, separated by underscores.
	for i := remainder; i < len(s); i += 3 {
		b.WriteByte('_')
		b.WriteString(s[i : i+3])
	}

	return b.String()
}
