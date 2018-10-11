package problem130

func solve(board [][]byte) {
	m := len(board)
	if m <= 2 {
		return
	}

	n := len(board[0])
	if n <= 2 {
		return
	}

	// 记录不能翻转的 'O'
	infection := make([][]bool, m)
	for k := range infection {
		infection[k] = make([]bool, n)
	}

	// 待检查的 i, j
	is := make([]int, 0, m*n)
	js := make([]int, 0, m*n)

	bfs := func(i, j int) {
		infection[i][j] = true
		is = append(is, i)
		js = append(js, j)

		for len(is) > 0 {
			i := is[0]
			j := js[0]
			is = is[1:]
			js = js[1:]

			// 检查该点的上下左后

			if i-1 >= 0 && board[i-1][j] == 'O' && !infection[i-1][j] {
				infection[i-1][j] = true
				is = append(is, i-1)
				js = append(js, j)
			}

			if i+1 < m && board[i+1][j] == 'O' && !infection[i+1][j] {
				infection[i+1][j] = true
				is = append(is, i+1)
				js = append(js, j)
			}

			if j-1 >= 0 && board[i][j-1] == 'O' && !infection[i][j-1] {
				infection[i][j-1] = true
				is = append(is, i)
				js = append(js, j-1)
			}

			if j+1 < n && board[i][j+1] == 'O' && !infection[i][j+1] {
				infection[i][j+1] = true
				is = append(is, i)
				js = append(js, j+1)
			}
		}

	}

	// 从四周的边开始检查

	for i := 0; i < m; i++ {
		if board[i][0] == 'O' && !infection[i][0] {
			bfs(i, 0)
		}

		if board[i][n-1] == 'O' && !infection[i][n-1] {
			bfs(i, n-1)
		}
	}

	for j := 0; j < n; j++ {
		if board[0][j] == 'O' && !infection[0][j] {
			bfs(0, j)
		}

		if board[m-1][j] == 'O' && !infection[m-1][j] {
			bfs(m-1, j)
		}
	}

	// 遍历
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' && !infection[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}
