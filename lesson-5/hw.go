package main

import (
	"fmt"
	"strings"
)

type ReturnStringType = func(string) string
type FunctionType func(op1 ReturnStringType) func(op2 ReturnStringType) string

func transformer(input string) FunctionType {
	return func(op1 ReturnStringType) func(op2 ReturnStringType) string {
		return func(op2 ReturnStringType) string {
			upperCased := op1(input)
			return op2(upperCased)
		}
	}
}

func toUpperCase(input string) string {
	return strings.ToUpper(input)
}
func reverseString(input string) string {
	reversed := []rune(input)
	n := len(reversed)
	for i := 0; i < n/2; i++ {
		reversed[i], reversed[n-i-1] = reversed[n-i-1], reversed[i]
	}

	return string(reversed)
}

func removeElement(nums []int, val int) int {
	p1, p2 := 0, len(nums)-1

	fmt.Println(nums, p2)
	for ; p1 <= p2; p1++ {
		if nums[p1] == val {
			nums[p1], nums[p2] = nums[p2], nums[p1]
			p2--
			p1--
		}
	}

	p2++

	fmt.Println(nums, p2)
	return p2
}

func main() {
	// op1 := transformer("hello")
	// op2 := op1(toUpperCase)
	// result := op2(reverseString)

	// fmt.Println(result, p1, p2)
	removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
}
