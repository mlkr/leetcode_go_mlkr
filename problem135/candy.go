package problem135

func candy(ratings []int) int {
	size := len(ratings)
	if size <= 1 {
		return size
	}

	left := make([]int, size)
	right := make([]int, size)

	left[0] = 1
	right[size-1] = 1

	for i := 1; i < size; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}

		if ratings[size-i-1] > ratings[size-i] {
			right[size-i-1] = right[size-i] + 1
		} else {
			right[size-i-1] = 1
		}
	}

	res := 0
	for i := 0; i < size; i++ {
		res += max(left[i], right[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
