package problem611

import "sort"

func triangleNumber(nums []int) int {
	size := len(nums)
	if size < 3 {
		return 0
	}
	sort.Ints(nums)

	res := 0
	for i := 2; i < size; i++ {
		if nums[i] >= nums[i-1]+nums[i-2] {
			continue
		}

		l, r := 0, i-1
		for l < r {
			if nums[l]+nums[r] > nums[i] {
				res += r - l
				r--
			} else {
				l++
			}
		}
	}

	return res
}
