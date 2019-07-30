package problem312

func maxCoins(nums []int) int {
	n := len(nums)
	newNums := make([]int, n+2)
	newNums[0], newNums[n+1] = 1, 1
	copy(newNums[1:n+1], nums)

	// dp[i][j] 是 nums[i:j+1] 的最大值
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}

	max := 0
	for i := n; i >= 1; i-- {
		for j := i; j <= n; j++ {
			for k := i; k <= j; k++ {
				coins := dp[i][k-1] + newNums[i-1]*newNums[k]*newNums[j+1] + dp[k+1][j]
				if coins > max {
					max = coins
				}
			}

			dp[i][j] = max
			max = 0
		}
	}

	return dp[1][n]
}
