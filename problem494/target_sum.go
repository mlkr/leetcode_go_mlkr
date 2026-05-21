package problem494

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	if target > sum {
		return 0
	}

	// sum+target 为偶数才有解
	if (sum+target)%2 != 0 {
		return 0
	}

	t := (sum + target) / 2
	if t < 0 {
		return 0
	}

	// sum(p) 为 nums 正数部分，sum(n) 为负数部分
	// sum(p) - sum(n) = target
	// sum(p) - sum(n) + sum(p) + sum(n) = target + sum(p) + sum(n)
	// 2*sum(p) = target + sum
	// sum(p) = (target + sum) / 2
	// 转化为在nums中和为sum(p)的排列组合数
	return getWays(nums, t)

}

func getWays(nums []int, target int) int {
	// dp[target] 为在nums中和为target的数的排列组合
	// dp[i] += do[i-n] 为通项公式
	// 参看例1
	dp := make([]int, target+1)
	dp[0] = 1

	for _, n := range nums {
		for i := target; i >= n; i-- {
			dp[i] += dp[i-n]
		}
	}

	return dp[target]
}
