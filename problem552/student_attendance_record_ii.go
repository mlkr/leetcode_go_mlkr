package problem552

func checkRecord(n int) int {
	switch n {
	case 1:
		return 3
	case 2:
		return 8
	case 3:
		return 19
	}

	dp := make([]int, n+1)
	// count without "A"
	// X    2
	// dp[0] = 2

	// XP    2
	// XL    2
	// dp[2] = 4

	// XXP    4
	// XPL    2
	// PLL    1
	// dp[3] = 7

	// XXXP    7 dp[i-1]
	// XXPL    4 dp[i-2]
	// XPLL    2 dp[i-3]
	// dp[4] = dp[3]+dp[2]+dp[1]
	dp[0] = 1
	dp[1] = 2
	dp[2] = 4
	dp[3] = 7
	for i := 4; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % 1000000007
	}

	// with "A"
	// AXXX
	// XAXX
	// XXAX
	// XXXA
	withA := 0
	for i := 0; i < n; i++ {
		withA = (withA + dp[i]*dp[n-1-i]) % 1000000007
	}

	return (dp[n] + withA) % 1000000007
}
