package main

import (
	"fmt"
	"sort"
)

// Homework-9
func main() {
	var numbers [3]int = [3]int{}
	for i := 0; i < 3; i++ {
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scan(&numbers[i])
	}
	sort.Ints(numbers[:])
	fmt.Printf("Sorted array: %v", numbers)
}
