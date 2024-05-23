package problem565

func arrayNesting(nums []int) int {
	res := 0
	for i := range nums {
		temp := 1
		for i != nums[i] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			temp++
		}

		res = max(res, temp)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
