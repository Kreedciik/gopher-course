package main

import "math"

type FuncResult func() int

func calculator(a, b int, operation string) FuncResult {

	return func() int {
		switch operation {
		case "+":
			return a + b
		case "-":
			return a - b
		case "/":
			return a / b
		case "*":
			return a * b
		case "%":
			return a % b
		default:
			return int(math.Pow(float64(a), float64(b)))
		}
	}
}

func main() {

}
