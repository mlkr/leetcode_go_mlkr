package problem198

func rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	if size == 1 {
		return nums[0]
	}

	mark := make([]int, size, size)
	mark[size-1] = nums[size-1]
	mark[size-2] = nums[size-2]

	for i := size - 1 - 2; i >= 0; i-- {
		if i+3 < size {
			mark[i] = nums[i] + max(mark[i+2], mark[i+3])
			continue
		}

		mark[i] = nums[i] + mark[i+2]
	}

	return max(mark[0], mark[1])
}

// 解法二
func rob2(nums []int) int {
	even, odd := 0, 0
	for k, v := range nums {
		if k%2 == 0 {
			even = max(even+v, odd)
		} else {
			odd = max(even, odd+v)
		}
	}

	return max(even, odd)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
