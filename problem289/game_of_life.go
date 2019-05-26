package problem289

func gameOfLife(board [][]int) {
	m := len(board)
	if m == 0 {
		return
	}

	n := len(board[0])
	if n == 0 {
		return
	}

	check := func(i, j int) {
		count := 0

		for row := i - 1; row <= i+1; row++ {
			for col := j - 1; col <= j+1; col++ {
				if 0 <= row && row < m && 0 <= col && col < n && !(row == i && col == j) {
					if board[row][col] == 1 || board[row][col] == 2 {
						count++
					}
				}
			}
		}

		if board[i][j] == 1 && (count < 2 || count > 3) {
			board[i][j] = 2
		}

		if board[i][j] == 0 && count == 3 {
			board[i][j] = 3
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			check(i, j)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] %= 2
		}
	}
}
