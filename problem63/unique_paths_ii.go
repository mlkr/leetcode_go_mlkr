package problem63

// func uniquePathsWithObstacles(obstacleGrid [][]int) int {

// 	// 用-1替代1, 表示此路不通
// 	for rowKey, row := range obstacleGrid {
// 		for colKey := range row {
// 			if obstacleGrid[rowKey][colKey] == 1 {
// 				obstacleGrid[rowKey][colKey] = -1
// 			}
// 		}
// 	}

// 	// 计算从原点开始直下直右的路劲数
// 	x := len(obstacleGrid)
// 	y := len(obstacleGrid[0])

// 	if obstacleGrid[0][0] != -1 {
// 		obstacleGrid[0][0] = 1
// 	}

// 	for i := 1; i < x; i++ {
// 		if obstacleGrid[i-1][0] == -1 || obstacleGrid[i][0] == -1 {
// 			obstacleGrid[i][0] = -1
// 		} else {
// 			obstacleGrid[i][0] = 1
// 		}
// 	}

// 	for j := 1; j < y; j++ {
// 		if obstacleGrid[0][j-1] == -1 || obstacleGrid[0][j] == -1 {
// 			obstacleGrid[0][j] = -1
// 		} else {
// 			obstacleGrid[0][j] = 1
// 		}
// 	}

// 	for i := 1; i < x; i++ {
// 		for j := 1; j < y; j++ {
// 			down := obstacleGrid[i-1][j]
// 			if down == -1 {
// 				down = 0
// 			}

// 			right := obstacleGrid[i][j-1]
// 			if right == -1 {
// 				right = 0
// 			}

// 			if obstacleGrid[i][j] != -1 {
// 				obstacleGrid[i][j] = down + right

// 			}
// 		}
// 	}

// 	if obstacleGrid[x-1][y-1] == -1 {
// 		obstacleGrid[x-1][y-1] = 0
// 	}

// 	return obstacleGrid[x-1][y-1]
// }

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}

	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}

	res := make([][]int, m)
	for r := range res {
		res[r] = make([]int, n)
	}

	if obstacleGrid[0][0] == 0 {
		res[0][0] = 1
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			res[i][0] = res[i-1][0]
		}
	}

	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			res[0][j] = res[0][j-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				res[i][j] = res[i-1][j] + res[i][j-1]
			}
		}
	}

	return res[m-1][n-1]
}
