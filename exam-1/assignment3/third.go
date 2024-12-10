package thirdassignment

func SearchMaxValue(slc []int) int {
	max := slc[0]

	for i := 1; i < len(slc); i++ {
		value := slc[i]
		if max < value {
			max = value
		}
	}

	return max
}
