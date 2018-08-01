package problem62

func uniquePaths(m int, n int) int {
	// res 存放到达坐标 i,j 的路径数
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	// 到达这些点的路径都为1
	for i := 0; i < m; i++ {
		res[i][0] = 1
	}

	for j := 0; j < n; j++ {
		res[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 到达该点的路劲数为, 上方点的路劲 + 左方点的路劲
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}

	return res[m-1][n-1]
}
