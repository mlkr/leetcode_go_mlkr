package problem91

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// 表示到第n个字符,有多少解法 s[:n+1]
	dp := make([]int, n+1)

	dp[0], dp[1] = 1, one(s[0:1])

	for i := 2; i <= n; i++ {
		// 分别计算当前1个字符, 2个字符是否合法
		w1, w2 := one(s[i-1:i]), two(s[i-2:i])

		// 与之前计算的结果结合
		dp[i] = w1*dp[i-1] + w2*dp[i-2]
		if dp[i] == 0 {
			dp[n] = 0
			break
		}
	}

	return dp[n]
}

// 表示一个字符是否合法
// 1-合法
// 0-不合法
func one(s string) int {
	if s == "0" {
		return 0
	}

	return 1
}

// 表示两个字符是否合法
// 1-合法
// 0-不合法
func two(s string) int {
	switch s[0] {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		if s[1] == '7' || s[1] == '8' || s[1] == '9' {
			return 0
		}

		return 1
	default:
		return 0
	}
}
