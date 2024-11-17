package main

import "fmt"

// Homework-3
func main() {

	var input string

	for {
		fmt.Print("Enter single character: ")
		fmt.Scan(&input)
		if len(input) == 1 {
			break
		} else {
			fmt.Println("Please enter single character!")
		}
	}

	char := rune(input[0])
	switch {
	case char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z':
		fmt.Println("Letter")
	case char >= '0' && char <= '9':
		fmt.Println("Digit")
	default:
		fmt.Println("Other symbol")
	}
}
