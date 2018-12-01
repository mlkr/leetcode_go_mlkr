package problem174

// 例子
// -2 -3 3
// -5 -10 1
// 10 30 -5

//  7 5 2
// 6 11 5
// 1 1  6
func calculateMinimumHP(dungeon [][]int) int {
	rows := len(dungeon)
	cols := len(dungeon[0])

	dungeon[rows-1][cols-1] = 1 - dungeon[rows-1][cols-1]
	if dungeon[rows-1][cols-1] <= 0 {
		dungeon[rows-1][cols-1] = 1
	}

	// 计算最后一行所需的生命值
	for c := cols - 2; c >= 0; c-- {
		if dungeon[rows-1][c] >= dungeon[rows-1][c+1] {
			dungeon[rows-1][c] = 1
		} else {
			dungeon[rows-1][c] = dungeon[rows-1][c+1] - dungeon[rows-1][c]
		}
	}

	// 计算最后一列所需的生命值
	for r := rows - 2; r >= 0; r-- {
		if dungeon[r][cols-1] >= dungeon[r+1][cols-1] {
			dungeon[r][cols-1] = 1
		} else {
			dungeon[r][cols-1] = dungeon[r+1][cols-1] - dungeon[r][cols-1]
		}
	}

	// 计算余下的生命值
	for r := rows - 2; r >= 0; r-- {
		for c := cols - 2; c >= 0; c-- {
			min := dungeon[r][c+1]
			if dungeon[r+1][c] < dungeon[r][c+1] {
				min = dungeon[r+1][c]
			}

			if dungeon[r][c] >= min {
				dungeon[r][c] = 1
			} else {
				dungeon[r][c] = min - dungeon[r][c]
			}
		}
	}

	return dungeon[0][0]
}

// 解法二
func calculateMinimumHP2(dungeon [][]int) int {
	rows := len(dungeon)
	cols := len(dungeon[0])

	intMax := 1<<63 - 1
	dp := make([][]int, rows+1)
	for r := range dp {
		dp[r] = make([]int, cols+1)
		for c := range dp[r] {
			dp[r][c] = intMax
		}
	}

	dp[rows][cols-1] = 1
	for r := rows - 1; r >= 0; r-- {
		for c := cols - 1; c >= 0; c-- {
			health := min(dp[r+1][c], dp[r][c+1]) - dungeon[r][c]
			dp[r][c] = max(health, 1)
		}
	}

	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
