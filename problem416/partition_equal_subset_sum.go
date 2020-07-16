package problem416

import "sort"

// 参看
// https://www.cnblogs.com/grandyang/p/5951422.html
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum&1 == 1 {
		return false
	}

	target := sum >> 1
	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}

		if dp[target] {
			break
		}
	}

	return dp[target]
}

// 更快
func canPartition2(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum&1 == 1 {
		return false
	}

	target := sum >> 1

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	var slove func(int, int) bool
	slove = func(i, sum int) bool {
		if sum == target {
			return true
		}

		if i >= len(nums) || nums[i] > target || sum > target {
			return false
		}

		return slove(i+1, sum+nums[i]) || slove(i+1, sum)
	}

	return slove(0, 0)
}
