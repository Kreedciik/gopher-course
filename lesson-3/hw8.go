package main

import (
	"fmt"
	"math/rand"
)

// Homework-8
func main() {
	var input int
	n := rand.Intn(10)
	fmt.Println("Game started!")
	for {
		fmt.Printf("Guess the number between [0, 9] ")
		fmt.Scan(&input)
		if n == input {
			fmt.Println("You found the number!")
			break
		} else {
			fmt.Println("Try one more time")
		}
	}
}
