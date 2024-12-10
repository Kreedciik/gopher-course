package fourthassignment

func FindMissingValue(arr []int) int {

	for i, v := range arr {
		if i+1 != v {
			return i + 1
		}
	}

	return 0
}
