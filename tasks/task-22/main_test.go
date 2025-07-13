package main

import (
	"math/big"
	"testing"
)

func TestAdd(t *testing.T) {
	a := big.NewInt(1000)
	b := big.NewInt(2000)
	expected := big.NewInt(3000)

	if result := add(a, b); result.Cmp(expected) != 0 {
		t.Errorf("add(%s, %s) = %s; want %s", a, b, result, expected)
	}
}

func TestSubtract(t *testing.T) {
	a := big.NewInt(5000)
	b := big.NewInt(3000)
	expected := big.NewInt(2000)

	if result := subtract(a, b); result.Cmp(expected) != 0 {
		t.Errorf("subtract(%s, %s) = %s; want %s", a, b, result, expected)
	}
}

func TestMultiply(t *testing.T) {
	a := big.NewInt(123)
	b := big.NewInt(456)
	expected := big.NewInt(56088)

	if result := multiply(a, b); result.Cmp(expected) != 0 {
		t.Errorf("multiply(%s, %s) = %s; want %s", a, b, result, expected)
	}
}

func TestDivide(t *testing.T) {
	a := big.NewInt(100)
	b := big.NewInt(4)
	expected := big.NewInt(25)

	if result := divide(a, b); result.Cmp(expected) != 0 {
		t.Errorf("divide(%s, %s) = %s; want %s", a, b, result, expected)
	}
}

func TestDivideByZeroPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("divide by zero did not panic")
		}
	}()

	a := big.NewInt(100)
	b := big.NewInt(0)
	_ = divide(a, b) // must call panic
}

func TestFormatNumber(t *testing.T) {
	cases := []struct {
		input    *big.Int
		expected string
	}{
		{big.NewInt(123), "123"},
		{big.NewInt(1234), "1_234"},
		{big.NewInt(1234567), "1_234_567"},
		{big.NewInt(9876543210), "9_876_543_210"},
	}

	for _, c := range cases {
		result := formatNumber(c.input)
		if result != c.expected {
			t.Errorf("formatNumber(%s) = %s; want %s", c.input, result, c.expected)
		}
	}
}
