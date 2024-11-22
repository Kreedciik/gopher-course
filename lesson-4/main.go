package main

import "math"

func containsDuplicate(arr []int) bool {
	m := map[int]bool{}

	for _, v := range arr {
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = true
		}
	}

	return false
}

func containsNearByDuplicate(arr []int, k int) bool {
	m := map[int]int{}

	for i, v := range arr {
		if _, ok := m[v]; ok && math.Abs(float64(i-m[v])) <= float64(k) {
			return true
		} else {
			m[v] = i
		}
	}

	return false
}
func main() {
	print(containsNearByDuplicate([]int{1, 2, 3, 4}, 2))

}

func calculateOddAndEven(a []int) ([]int, []int) {
	odd, even := []int{}, []int{}
	for _, v := range a {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	return odd, even
}

func duplicateZeros(arr []int) {
	for i, v := range arr {
		if v == 0 {
			arr = append(arr[:i+1], append([]int{0}, arr[i+1:len(arr)-1]...)...)
		}
	}
}
