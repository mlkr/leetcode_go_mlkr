package problem446

// 先看 413 题
func numberOfArithmeticSlices(A []int) int {
	size := len(A)
	if size < 3 {
		return 0
	}

	dp := make([]map[int]int, size)
	res := 0

	min, max := -1<<31, 1<<31-1

	for i := 1; i < size; i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			diff := A[i] - A[j]

			if diff >= max || diff <= min {
				continue
			}

			dp[i][diff] += dp[j][diff] + 1
			res += dp[j][diff]
		}
	}

	return res
}
