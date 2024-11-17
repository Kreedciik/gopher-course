package main

import (
	"fmt"
	"math"
	"strings"
)

// Homework-10
func main() {
	var numbers = [2]int{}
	var operation string
	mathOperations := "+-/*%^"

	fmt.Println("Simple calculator")
	fmt.Println()

	for i := 0; i < 2; i++ {
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	for {
		fmt.Printf("Enter math operation (%v): ", mathOperations)
		fmt.Scan(&operation)
		if strings.Contains(mathOperations, operation) && len(operation) == 1 {
			break
		} else {
			fmt.Println("Invalid math operation")
		}
	}

	a, b := numbers[0], numbers[1]
	switch operation {
	case "+":
		fmt.Printf("%d + %d = %d", a, b, a+b)
	case "-":
		fmt.Printf("%d - %d = %d", a, b, a-b)
	case "/":
		fmt.Printf("%d / %d = %d", a, b, a/b)
	case "*":
		fmt.Printf("%d * %d = %d", a, b, a*b)
	case "%":
		fmt.Printf("%d %% %d = %d", a, b, a%b)
	default:
		fmt.Printf("%d ^ %d = %v", a, b, math.Pow(float64(a), float64(b)))
	}
}
