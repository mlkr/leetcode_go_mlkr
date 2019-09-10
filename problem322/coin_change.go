package problem322

func coinChange(coins []int, amount int) int {
	// dp[i] 更换金额为 i 时，最小的零钱数量
	// res = dp[amount]
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, c := range coins {
			if i-c >= 0 && dp[i] > dp[i-c]+1 {
				dp[i] = dp[i-c] + 1
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
