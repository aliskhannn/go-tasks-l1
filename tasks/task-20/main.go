package main

import (
	"fmt"
)

// function to reverse the whole string
func reverse(b []byte, start, end int) {
	for start < end {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
}

func reverseWords(s string) string {
	b := []byte(s) // convert string to slice of bytes

	// reverse the whole string
	reverse(b, 0, len(b)-1) // nus god wons

	// then reverse each word back separately by space
	// like this -> sun god wons
	//	 		 ->	sun dog wons
	//	  		 ->	sun dog snow
	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b, start, i-1)
			start = i + 1
		}
	}

	return string(b)
}

func main() {
	s := "snow dog sun"

	fmt.Println(reverseWords(s)) // sun dog snow
}
