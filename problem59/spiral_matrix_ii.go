package problem59

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	res := make([][]int, n, n)

	for i := range res {
		res[i] = make([]int, n, n)
	}

	top, right, bottom, left := 0, n-1, n-1, 0
	num := 1
	for top <= bottom && left <= right {
		// →
		for j := left; j <= right; j++ {
			res[top][j] = num
			num++
		}
		top++

		// ↓
		for j := top; j <= bottom; j++ {
			res[j][right] = num
			num++
		}
		right--

		// ←
		for j := right; j >= left; j-- {
			res[bottom][j] = num
			num++
		}
		bottom--

		// ↑
		for j := bottom; j >= top; j-- {
			res[j][left] = num
			num++
		}
		left++
	}

	return res
}
