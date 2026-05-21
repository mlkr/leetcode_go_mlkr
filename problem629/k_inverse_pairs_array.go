package problem629

const m = 1000000007

// 参看https://www.cnblogs.com/kexinxin/p/10381439.html
func kInversePairs(n int, k int) int {
	if k == 0 {
		return 1
	}

	if n == 1 {
		return 0
	}

	// dp[n][k] 1-n k个逆序对数
	// xxxx5 dp[n][k]
	// xxx5x dp[n][k-1]
	// xx5xx dp[n][k-2]
	// x5xxx dp[n][k-3]
	// 5xxxx dp[n][k-4]
	// dp[n+1][k] = dp[n][k]+dp[n][k-1]+...+dp[n][k-n]
	// dp[n][k] = dp[n-1][k]+dp[n-1][k-1]+...+dp[n-1][k-n+1] （1）
	// dp[n][k+1] = dp[n-1][k+1]+dp[n-1][k]+dp[n-1][k-1]+...+dp[n-1][k+1-n+1] （2）
	// dp[n][k+1]-dp[n][k] = dp[n-1][k+1]-dp[n-1][k-n+1]  (2)-(1)
	// dp[n][k+1] = dp[n][k]+dp[n-1][k+1]-dp[n-1][k-n+1]

	// do[n][k] = dp[n][k-1]+db[n-1][k]-dp[n-1][k-n]
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		dp[i][0] = 1
	}

	for i := 2; i <= n; i++ {
		// 1-n 最多有 n*(n-1)/2 个逆序对
		// 4321
		// (n-1)+(n-2)+...+1
		// (n-1+1)/2*(n-1)
		// n*(n-1)/2
		maxJ := min(k, i*(i-1)/2)
		for j := 1; j <= maxJ; j++ {
			dp[i][j] = (dp[i][j-1] + dp[i-1][j]) % m
			if j >= i {
				dp[i][j] -= dp[i-1][j-i]
				if dp[i][j] < 0 {
					dp[i][j] += m
				}
			}
		}

	}

	return dp[n][k]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
