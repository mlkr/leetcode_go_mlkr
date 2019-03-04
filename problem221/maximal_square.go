package problem221

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}

	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	dp := make([][]int, m)
	maxEdge := 0

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxEdge = 1
		}
	}

	for j := 1; j < n; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			maxEdge = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1 +
					min(
						dp[i-1][j-1],
						min(
							dp[i-1][j],
							dp[i][j-1],
						),
					)
				maxEdge = max(maxEdge, dp[i][j])
			}
		}
	}

	return maxEdge * maxEdge
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
