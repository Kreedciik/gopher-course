package main

import (
	"fmt"
	"math/rand"
)

// Homework-5
func main() {
	a := make([]int, 10)
	b := make([]int, 10)
	var matched = []int{}

	for i := 0; i < len(a); i++ {
		randInt1 := rand.Intn(100)
		randInt2 := rand.Intn(100)
		a[i] = randInt1
		b[i] = randInt2
		if randInt1 == randInt2 {
			matched = append(matched, randInt1)
		}
	}
	fmt.Println("Slices: ", a, b)

	if len(matched) > 0 {
		fmt.Print("Matched elements: ")
		for i := 0; i < len(matched); i++ {
			fmt.Printf(" %v", matched[i])
		}
	} else {
		fmt.Println("No matches!")
	}
}
