package problem300

import "sort"

func lengthOfLIS(nums []int) int {
	size := len(nums)
	increment := make([]int, 0, size)
	for i := 0; i < size; i++ {
		at := sort.SearchInts(increment, nums[i])
		if at == len(increment) {
			increment = append(increment, nums[i])
		} else {
			increment[at] = nums[i]
		}
	}

	return len(increment)
}
