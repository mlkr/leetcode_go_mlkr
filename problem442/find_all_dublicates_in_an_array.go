package problem442

func findDuplicates(nums []int) []int {
	res := []int{}

	for _, num := range nums {
		idx := abs(num) - 1
		if nums[idx] < 0 {
			res = append(res, idx+1)
		}

		nums[idx] *= -1
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func findDuplicates2(nums []int) []int {
	res := []int{}

	for _, num := range nums {
		idx := (num & (1<<30 - 1)) - 1
		if nums[idx]&(1<<30) > 0 {
			res = append(res, idx+1)
		}

		nums[idx] |= 1 << 30
	}

	return res
}
