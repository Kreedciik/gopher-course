package main

import (
	"fmt"
	"strings"
)

// Homework-2
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

	isUpperCase := strings.ToUpper(input) == input

	if isUpperCase {
		fmt.Println("Uppercase")
	} else {
		fmt.Println("Lowercase")
	}

}
