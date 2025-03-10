package main

import (
	"fmt"
	"strings"
)

func countLetters(word string, resultCh chan map[rune]int) {
	result := make(map[rune]int)

	word = strings.ToLower(word)
	for _, chr := range word {
		if chr >= 'a' && chr <= 'z' {
			result[chr]++
		}
	}
	resultCh <- result
}

func main() {
	var strArr []string = []string{"quick", "brown", "fox", "lazy", "dog"}
	resultCh := make(chan map[rune]int)

	for _, word := range strArr {
		go countLetters(word, resultCh)
	}

	finalCount := make(map[rune]int)
	for i := 0; i < len(strArr); i++ {
		wordCount := <-resultCh
		for letter, count := range wordCount {
			finalCount[letter] += count
		}
	}

	fmt.Println("{")
	for char := 'a'; char <= 'z'; char++ {
		fmt.Printf("    \"%c\": %d,\n", char, finalCount[char])
	}
	fmt.Println("}")
}
