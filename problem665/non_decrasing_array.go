package problem665

import "sort"

func checkPossibility(nums []int) bool {
	size := len(nums)
	for i := 1; i < size; i++ {
		if nums[i-1] > nums[i] {
			pre := deepCopy(nums)
			pre[i] = pre[i-1]

			next := deepCopy(nums)
			next[i-1] = next[i]

			return sort.IsSorted(sort.IntSlice(pre)) || sort.IsSorted(sort.IntSlice(next))
		}
	}

	return true
}

func deepCopy(nums []int) []int {
	numsC := make([]int, len(nums))
	copy(numsC, nums)

	return numsC
}
