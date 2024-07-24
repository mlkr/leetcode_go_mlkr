package problem583

func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	// dp[i][j] 长度为i的word1,长度为j的word2相同的字符数
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}

	return n1 + n2 - dp[n1][n2]*2
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
