package problem628

func maximumProduct(nums []int) int {
	max, max1, max2 := -1001, -1001, -1001
	min, min1 := 1001, 1001

	for _, n := range nums {
		switch {
		case n > max:
			max2 = max1
			max1 = max
			max = n
		case n > max1:
			max2 = max1
			max1 = n
		case n > max2:
			max2 = n
		}

		switch {
		case n < min:
			min1 = min
			min = n
		case n < min1:
			min1 = n
		}
	}

	return bigger(max*max1*max2, max*min1*min)
}

func bigger(a, b int) int {
	if a > b {
		return a
	}

	return b
}
