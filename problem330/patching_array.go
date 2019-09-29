package problem330

// https://leetcode.com/problems/patching-array/discuss/78490/Share-my-greedy-solution-by-Java-with-simple-explanation-(time%3A-1-ms)
func minPatches(nums []int, n int) int {

	count, i, max := 0, 0, 0
	for max < n {
		if i >= len(nums) || nums[i]-1 > max {
			max += max + 1
			count++
		} else {
			// max: [1, max] 区间内的任何数字都可以用 nums[:i+1] 中某几个数字的和表示
			max += nums[i]
			i++
		}
	}

	return count
}
