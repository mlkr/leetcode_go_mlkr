package problem120

func minimumTotal(triangle [][]int) int {
	tLen := len(triangle)
	if tLen == 0 {
		return 0
	}

	if tLen == 1 {
		return triangle[0][0]
	}

	for i := 1; i < tLen; i++ {
		for j := range triangle[i] {
			switch {
			case j == 0:
				triangle[i][j] += triangle[i-1][j]
			case j == i:
				triangle[i][j] += triangle[i-1][j-1]
			case triangle[i-1][j-1] < triangle[i-1][j]:
				triangle[i][j] += triangle[i-1][j-1]
			default:
				triangle[i][j] += triangle[i-1][j]
			}
		}
	}

	min := triangle[tLen-1][0]
	for i := 1; i < len(triangle[tLen-1]); i++ {
		if triangle[tLen-1][i] < min {
			min = triangle[tLen-1][i]
		}
	}

	return min
}

// 另一解法
// 从下往上递归处理
//   2       10
//  3 4     8  10
// 5 6 7   5  6  7
func minimumTotal2(triangle [][]int) int {
	tLen := len(triangle)
	if tLen == 0 {
		return 0
	}

	if tLen == 1 {
		return triangle[0][0]
	}

	dp := make([]int, len(triangle[tLen-1]))
	copy(dp, triangle[tLen-1])

	for i := tLen - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j]+triangle[i][j], dp[j+1]+triangle[i][j])
		}
	}

	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
