package problem279

import "math"

func numSquares(n int) int {
	// 四平方和定理
	// http://www.cnblogs.com/grandyang/p/4800552.html
	for n%4 == 0 {
		n /= 4
	}

	if n%8 == 7 {
		return 4
	}

	res := 3
	for i := 1; i*i <= n; i++ {
		j := math.Sqrt(float64(n - i*i))
		if int(j) == 0 {
			res = 1
			break
		}

		if int(j)*int(j)+i*i == n {
			res = 2
		}
	}

	return res
}

// 解法二
func numSquares2(n int) int {
	squares := []int{}
	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}

	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	for _, p := range squares {
		for i := p; i <= n; i++ {
			if dp[i] > dp[i-p]+1 {
				dp[i] = dp[i-p] + 1
			}
		}
	}

	return dp[n]
}
