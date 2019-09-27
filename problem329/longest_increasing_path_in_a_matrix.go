package problem329

// 上下左右移动
var dirs = [][]int{
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
	[]int{-1, 0},
}

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}

	n := len(matrix[0])

	cached := make([][]int, m)
	for i := range cached {
		cached[i] = make([]int, n)
	}

	maxLen := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			len := dfs(matrix, m, n, i, j, cached)
			maxLen = max(maxLen, len)
		}
	}

	return maxLen
}

func dfs(matrix [][]int, m, n, i, j int, cached [][]int) int {
	if cached[i][j] != 0 {
		return cached[i][j]
	}

	maxLen := 1
	for d := range dirs {
		x := i + dirs[d][0]
		y := j + dirs[d][1]

		if x >= m || x < 0 || y >= n || y < 0 || matrix[x][y] <= matrix[i][j] {
			continue
		}

		len := 1 + dfs(matrix, m, n, x, y, cached)
		maxLen = max(maxLen, len)
	}

	cached[i][j] = maxLen

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
