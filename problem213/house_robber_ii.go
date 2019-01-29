package problem213

func rob(nums []int) int {
	size := len(nums)
	switch size {
	case 0:
		return 0
	case 1:
		return nums[0]
	default:
		// 首尾不能相连
		return max(robbing(nums[1:]), robbing(nums[:size-1]))
	}
}

func robbing(nums []int) int {
	odd, even := 0, 0
	for k, v := range nums {
		if k%2 != 0 {
			odd = max(odd+v, even)
		} else {
			even = max(odd, even+v)
		}
	}

	return max(odd, even)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
