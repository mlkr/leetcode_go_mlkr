package problem115

// 使用动态规划
// 依次匹配 r, ra,..., rabbit
// 	    r a b b b i t
// r	1 1 1 1 1 1 1
// a	0 1 1 1 1 1 1
// b	0 0 1 2 3 3 3
// b	0 0 0 1 3 3 3
// i	0 0 0 0 0 3 3
// t	0 0 0 0 0 0 3

// 		b a b g b a g
// b	1 1 2 2 3 3 3
// a	0 1 1 1 1 4 4
// g	0 0 0 1 1 1 5

func numDistinct(s string, t string) int {
	sLen, tLen := len(s), len(t)
	dp := make([][]int, tLen+1)
	for k := range dp {
		dp[k] = make([]int, sLen+1)
	}

	for si := 0; si <= sLen; si++ {
		dp[0][si] = 1
	}

	for ti := 1; ti <= tLen; ti++ {
		for si := ti; si <= sLen; si++ {
			dp[ti][si] = dp[ti][si-1]
			if s[si-1] == t[ti-1] {
				dp[ti][si] += dp[ti-1][si-1]
			}
		}
	}

	return dp[tLen][sLen]
}

// 另一解法(与上面的解法原理一样, 但是简化 dp 为一维数组)(最佳)
func numDistinct2(s string, t string) int {
	sLen, tLen := len(s), len(t)
	dp := make([]int, sLen+1)
	for k := range dp {
		dp[k] = 1
	}

	var pre int
	for ti := 1; ti <= tLen; ti++ {
		pre, dp[ti-1] = dp[ti-1], 0
		for si := ti; si <= sLen; si++ {
			if s[si-1] == t[ti-1] {
				pre, dp[si] = dp[si], dp[si-1]+pre
			} else {
				pre, dp[si] = dp[si], dp[si-1]
			}
		}
	}

	return dp[sLen]
}
