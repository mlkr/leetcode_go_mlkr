package problem456

func find132pattern(nums []int) bool {
	stack := make([]int, 0, len(nums))
	r := -1 << 31
	for l := len(nums) - 1; l >= 0; l-- {
		if nums[l] < r {
			return true
		}

		for len(stack) > 0 && nums[l] > stack[len(stack)-1] {
			r = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, nums[l])
	}

	return false
}
