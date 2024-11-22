package main

import (
	"fmt"
	"sort"
)

func main() {
	var num1, num2, num3 int

	fmt.Printf("Enter number %d: ", 1)
	fmt.Scan(&num1)

	fmt.Printf("Enter number %d: ", 2)
	fmt.Scan(&num2)

	fmt.Printf("Enter number %d: ", 3)
	fmt.Scan(&num3)

	numArr := []int{num1, num2, num3}
	sort.Ints(numArr)

	fmt.Println("Sorted numbers: ", numArr)
}
