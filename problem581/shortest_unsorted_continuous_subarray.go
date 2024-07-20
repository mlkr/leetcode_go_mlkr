package problem581

func findUnsortedSubarray(nums []int) int {
	size := len(nums)
	if size == 1 {
		return 0
	}

	left, right := 1, 0
	min, max := nums[size-1], nums[0]

	for i := 1; i < size; i++ {
		if nums[i] >= max {
			max = nums[i]
		} else {
			right = i
		}

		j := size - 1 - i
		if nums[j] <= min {
			min = nums[j]
		} else {
			left = j
		}
	}

	return right - left + 1
}

// 2 6 4 15 10 8 9
