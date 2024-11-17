package main

import (
	"fmt"
	"strings"
)

// Homework-7
func main() {
	var text string

	fmt.Print("Enter your text: ")
	fmt.Scan(&text)

	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		word := words[i]
		runes := []rune(word)
		n := len(runes)
		for j := 0; j < n/2; j++ {
			runes[j], runes[n-1-j] = runes[n-1-j], runes[j]
		}
		words[i] = string(runes)
	}

	fmt.Printf("The reversed text: %s", strings.Join(words, " "))
}
