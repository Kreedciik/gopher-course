package main

import (
	"fmt"
	"math/rand"
)

// Homework-6
func main() {
	a := make([]int, 20)
	k, count := 10, 0

	for i := 0; i < len(a); i++ {
		randInt := rand.Intn(100)
		a[i] = randInt
		if randInt < k {
			count++
		}
	}

	fmt.Println("Slice: ", a)
	fmt.Printf("The number of less than %v: %v", k, count)
}
