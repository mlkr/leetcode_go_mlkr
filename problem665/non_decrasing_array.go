package problem665

import "sort"

func checkPossibility(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			pre := nums[i-1]
			next := nums[i]

			nums[i-1] = next
			if sort.IntsAreSorted(nums) {
				return true
			}

			nums[i-1], nums[i] = pre, pre
			if sort.IntsAreSorted(nums) {
				return true
			}

			return false
		}
	}

	return true
}
