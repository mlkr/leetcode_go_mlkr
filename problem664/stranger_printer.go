package problem664

func strangePrinter(s string) int {
	s = removeDup(s)

	size := len(s)
	if size <= 1 {
		return size
	}

	// dp[i][j] 表示s[i:j+1] 最小打印次
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		dp[i][i] = 1
	}

	for Len := 2; Len <= size; Len++ {
		for i := 0; i+Len-1 < size; i++ {
			j := i + Len - 1
			dp[i][j] = dp[i][j-1] + 1

			// 新增的字符与最后一个字符相同
			if s[j] == s[j-1] {
				dp[i][j]--
			}

			// 原来的字符串中某个字符与新增的字符相同
			for k := i; k < j; k++ {
				if s[k] == s[j] {
					dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j]-1)
				}
			}
		}
	}

	return dp[0][size-1]
}

// 去掉重复的字符，因为重复字符只打印一次
func removeDup(s string) string {
	res := make([]byte, 1, len(s))
	res[0] = s[0]

	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			res = append(res, s[i])
		}
	}

	return string(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
