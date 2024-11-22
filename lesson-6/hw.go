package main

import "fmt"

// Homework-1
func swap(a, b *int) {
	*a, *b = *b, *a
}

// Homework-2
func keepNumbers(str *string) {
	numbers := ""
	for i := 0; i < len(*str); i++ {
		char := (*str)[i]
		if char >= '0' && char <= '9' {
			numbers += string(char)
		}
	}
	*str = numbers
}

// Homework-3
func increaseLengthByDouble(slc *[]int) {
	currentLength := len(*slc)
	for i := currentLength; i < 2*currentLength; i++ {
		*slc = append(*slc, (*slc)[i-currentLength])
	}
}

// Homework-4
// Apply bubble sort by ASC
func sortSlice(slc *[]int) {
	n := len(*slc)
	swapped := false
	for i := 0; i < n; i++ {
		swapped = false
		for j := 0; j < n-i-1; j++ {
			if (*slc)[j] > (*slc)[j+1] {
				temp := (*slc)[j]
				(*slc)[j] = (*slc)[j+1]
				(*slc)[j+1] = temp
				swapped = true
			}
		}

		if swapped == false {
			break
		}

	}
}

func main() {
	fmt.Println("Homework-1 result:")
	a, b := 1, 2
	fmt.Printf("Before swapping a=%d b=%d \n", a, b)
	swap(&a, &b)
	fmt.Printf("After swapping a=%d b=%d \n", a, b)
	fmt.Println("--------------------")

	fmt.Println("Homework-2 result:")
	str := "1hello34ar"
	fmt.Printf("Before str=%s \n", str)
	keepNumbers(&str)
	fmt.Printf("After str=%s \n", str)
	fmt.Println("--------------------")

	fmt.Println("Homework-3 result:")
	s := make([]int, 4)
	fmt.Printf("Before increasing: %v %v \n", len(s), s)
	increaseLengthByDouble(&s)
	fmt.Printf("After increasing: %v %v \n", len(s), s)
	fmt.Println("--------------------")

	fmt.Println("Homework-4 result:")
	slc := []int{3, 5, 3, 4, 1, 2}
	fmt.Printf("Before sorting: %v \n", slc)
	sortSlice(&slc)
	fmt.Printf("After sorting: %v \n", slc)
}
