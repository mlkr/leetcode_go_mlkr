package problem200

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])

	x := make([]int, 0, m*n)
	y := make([]int, 0, m*n)

	var add = func(i, j int) {
		x = append(x, i)
		y = append(y, j)
		grid[i][j] = '0'
	}

	var pop = func() (int, int) {
		i := x[0]
		x = x[1:]
		j := y[0]
		y = y[1:]

		return i, j
	}

	var bfs = func(i, j int) int {
		if grid[i][j] == '0' {
			return 0
		}

		add(i, j)
		for len(x) > 0 {
			i, j = pop()
			// 检查上下左右四个点
			if i-1 >= 0 && grid[i-1][j] == '1' {
				add(i-1, j)
			}

			if i+1 < m && grid[i+1][j] == '1' {
				add(i+1, j)
			}

			if j-1 >= 0 && grid[i][j-1] == '1' {
				add(i, j-1)
			}

			if j+1 < n && grid[i][j+1] == '1' {
				add(i, j+1)
			}
		}

		return 1
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res += bfs(i, j)
		}
	}

	return res
}
