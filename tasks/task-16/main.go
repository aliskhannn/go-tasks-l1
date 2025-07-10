package main

import "fmt"

// Быстрая сортировка (quicksort)
// Реализовать алгоритм быстрой сортировки массива встроенными средствами языка.
// Можно использовать рекурсию.
// Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел.
// Для выбора опорного элемента можно взять середину или первый элемент.

func quickSort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	pivot := s[0]           // selecting the first element as pivot
	var less, greater []int // slices for storing less and greater values

	for _, v := range s[1:] {
		if v <= pivot { // if v is less then pivot
			less = append(less, v) // append it to 'fewer' slice
		} else {
			greater = append(greater, v) // if greater append to 'greater' slice
		}
	}

	// recursively sort the array
	return append(append(quickSort(less), pivot), quickSort(greater)...)
}

func main() {
	s := []int{5, 3, 6, 8, 4, 9, 2, 7, 1, 10}

	fmt.Println(quickSort(s))
}
