package main

import (
	"fmt"
	"strings"
)

// Небольшой фрагмент кода — проблемы и решение.
// Рассмотреть следующий код и ответить на вопросы:
// к каким негативным последствиям он может привести и как это исправить?
//
// Приведите корректный пример реализации.

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
