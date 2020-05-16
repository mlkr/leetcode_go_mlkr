package problem377

import "sort"

func combinationSum4(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	// dp[i] 是 nums 中叠加等于 i 组合的数量
	// dp[target] 为答案
	dp := make([]int, target+1)
	dp[0] = 1

loop:
	for i := nums[0]; i <= target; i++ {
		for _, num := range nums {
			if i < num {
				continue loop
			}
			dp[i] += dp[i-num]
		}
	}

	return dp[target]
}
