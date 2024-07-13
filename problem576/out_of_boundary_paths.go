package problem576

const mod = 1000000007

func findPaths(m int, n int, maxMove int, startRow int, startCol int) int {
	size := m
	if n > m {
		size = n
	}

	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
	}

	for k := 0; k < maxMove; k++ {
		pre := make([]int, n)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				paths := 0

				if i == 0 {
					paths++
				} else {
					paths += pre[j]
				}

				if j == 0 {
					paths++
				} else {
					paths += pre[j-1]
				}

				if i == m-1 {
					paths++
				} else {
					paths += dp[i+1][j]
				}

				if j == n-1 {
					paths++
				} else {
					paths += dp[i][j+1]
				}

				pre[j] = dp[i][j]
				dp[i][j] = paths % mod
			}
		}
	}

	return dp[startRow][startCol]
}
