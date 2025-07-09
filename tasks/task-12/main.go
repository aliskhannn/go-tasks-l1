package main

import "fmt"

// Собственное множество строк.
// Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree").
// Создать для неё собственное множество.
//
// Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

func main() {
	m := make(map[string]bool)
	words := []string{"cat", "cat", "dog", "dog", "tree"}
	uniqueWords := make([]string, 0)

	for _, word := range words {
		if _, ok := m[word]; !ok { // if an element doesn't exist in m
			m[word] = true // mark as added
			uniqueWords = append(uniqueWords, word)
		}
	}

	fmt.Println(uniqueWords)
}
