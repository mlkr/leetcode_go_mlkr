package problem375

import "math"

// 保证能够赢得游戏且代价最小的方法是：
// 以最优策略在(1, n)的范围内选数字，
// 保证每一步都是错误的，直到最后一步为止
// (每次以最优方式选数字的方法，实际上是每个数字都试一下，看一下选哪个数字能够使得最后的结果最小就行)
// 参看
// https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/solution/cai-shu-zi-da-xiao-2-java-cong-bao-li-di-gui-dao-d/
// https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/solution/cai-shu-zi-da-xiao-ii-by-leetcode/
func getMoneyAmount(n int) int {
	if n < 2 {
		return 0
	}

	// dp[1][n] 为(1, n)范围内的最小花费
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(a, b int) int
	dfs = func(a, b int) int {
		if a >= b {
			return 0
		}

		if dp[a][b] != -1 {
			return dp[a][b]
		}

		mx := math.MaxInt32
		for i := (a + b) / 2; i <= b; i++ {
			mx = min(mx, i+max(dfs(a, i-1), dfs(i+1, b)))
		}

		dp[a][b] = mx

		return mx
	}

	return dfs(1, n)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
