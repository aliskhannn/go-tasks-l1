package main

import (
	"fmt"
	"strings"
)

var justString string

// example of createHugeString function
func createHugeString(size int) string {
	data := make([]byte, size)

	for i := 0; i < size; i++ {
		data[i] = 'A'
	}

	return string(data)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()

	fmt.Println(justString)
}
