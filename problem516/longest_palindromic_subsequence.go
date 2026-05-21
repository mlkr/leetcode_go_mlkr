package problem516

func longestPalindromeSubseq(s string) int {
	n := len(s)

	// dp[i][j] 为 s[i:j+1] 最长回文长度
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for size := 2; size <= n; size++ {
		for i := 0; i+size-1 < n; i++ {
			j := i + size - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
