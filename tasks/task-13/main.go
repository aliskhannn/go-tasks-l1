package main

import "fmt"

func main() {
	var a, b = 6, 14
	var c, d = 8, 23

	// example with addition/subtraction
	fmt.Printf("initial values: a = %d, b = %d\n", a, b)

	a = a + b // 6 + 14 = 20
	b = a - b // 20 - 14 = 6
	a = a - b // 20 - 6 = 14

	fmt.Printf("+/- result: a = %d, b = %d\n\n", a, b) // 14, 6

	// _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_ for dividing examples :) _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_

	// example with XOR
	fmt.Printf("initial values: c = %d, d = %d\n", c, d)

	c = c ^ d // 01000 ^ 10111 = 11111 - at this step c = 31
	d = c ^ d // 11111 ^ 10111 = 01000 - at this step d = 8
	c = c ^ d // 11111 ^ 01000 = 10111 - at this step c = 23

	fmt.Printf("XOR result: c = %d, d = %d\n", c, d) // 23, 8 | 10111, 01000
}
