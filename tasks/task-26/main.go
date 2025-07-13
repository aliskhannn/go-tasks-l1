package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	m := make(map[rune]bool)

	sLower := strings.ToLower(s)

	for _, r := range sLower {
		if m[r] {
			return false
		}
		m[r] = true
	}

	return true
}

func main() {
	s := "abcd"

	fmt.Println(isUnique(s))
}
