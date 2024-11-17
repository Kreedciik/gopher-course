package main

import (
	"fmt"
	"strings"
)

// Homework 1
func main() {
	vowels := "aeiouAEIOU"
	var input string

	for {
		fmt.Print("Enter the letter: ")
		fmt.Scan(&input)
		if len(input) == 1 {
			break
		} else {
			fmt.Println("Please enter single character!")
		}
	}

	if strings.Contains(vowels, input) {
		fmt.Println("Vowel")
	} else {
		fmt.Println("Not vowel")
	}

}
