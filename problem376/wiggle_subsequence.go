package problem376

func wiggleMaxLength(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}

	sign := nums[1] - nums[0]
	res := 2
	if sign == 0 {
		res = 1
	}

	var count func(diff int)
	count = func(diff int) {
		product := sign * diff

		if product < 0 || (sign == 0 && diff != 0) {
			res++
			sign = diff
		}

	}

	for i := 2; i < size; i++ {
		count(nums[i] - nums[i-1])
	}

	return res
}
