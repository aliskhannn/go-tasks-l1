package main

import "fmt"

func StringSet(words []string) []string {
	m := make(map[string]bool)
	uniqueWords := make([]string, 0)

	for _, word := range words {
		if !m[word] {
			m[word] = true
			uniqueWords = append(uniqueWords, word)
		}
	}

	return uniqueWords
}

func main() {
	words := []string{"cat", "cat", "dog", "dog", "tree"}
	fmt.Println(StringSet(words)) // [cat dog tree]
}
