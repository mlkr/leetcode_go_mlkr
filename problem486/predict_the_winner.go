package problem483

func PredictTheWinner(nums []int) bool {
	size := len(nums)

	// dp[i][j] 表示 nums[i] 到 nums[j] 范围 play1 比 play2 多出的分数
	// 相对的 -dp[i][j] 表示该范围内 play2 比 play1 多出的分数
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, size)
		dp[i][i] = nums[i]
	}

	for leng := 2; leng <= size; leng++ {
		for i := 0; i <= size-leng; i++ {
			j := i + leng - 1
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}

	return dp[0][size-1] >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
