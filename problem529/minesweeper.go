package problem529

var dx = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dy = []int{-1, 0, 1, -1, 1, -1, 0, 1}

func updateBoard(board [][]byte, click []int) [][]byte {
	x := click[0]
	y := click[1]

	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}

	return bfs(board, x, y)
}

func bfs(board [][]byte, x, y int) [][]byte {
	m := len(board)
	n := len(board[0])

	isHave := make([]bool, m*n)
	isHave[x*n+y] = true

	xs := make([]int, 1, m)
	ys := make([]int, 1, n)
	xs[0] = x
	ys[0] = y

	for len(xs) > 0 {
		x := xs[0]
		y := ys[0]
		xs = xs[1:]
		ys = ys[1:]

		bombs := getBombs(board, x, y)
		if bombs > 0 {
			board[x][y] = '0' + byte(bombs)
			continue
		}

		board[x][y] = 'B'
		for i := 0; i < 8; i++ {
			nextX := x + dx[i]
			nextY := y + dy[i]
			if nextX > -1 && nextX < m &&
				nextY > -1 && nextY < n &&
				board[nextX][nextY] == 'E' &&
				!isHave[nextX*n+nextY] {
				xs = append(xs, nextX)
				ys = append(ys, nextY)
				isHave[nextX*n+nextY] = true

			}
		}
	}

	return board
}

func getBombs(board [][]byte, x, y int) int {
	bombs := 0
	m, n := len(board), len(board[0])

	for i := 0; i < 8; i++ {
		nextX := x + dx[i]
		nextY := y + dy[i]
		if nextX > -1 && nextX < m &&
			nextY > -1 && nextY < n {
			if board[nextX][nextY] == 'M' {
				bombs++
			}
		}
	}

	return bombs
}

// dfs 更快
func updateBoard2(board [][]byte, click []int) [][]byte {
	x := click[0]
	y := click[1]

	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}

	m := len(board)
	n := len(board[0])
	had := make(map[int]bool, m*n)

	var dfs func(x, y int)
	dfs = func(x, y int) {
		had[x*n+y] = true

		bombs := getBombs(board, x, y)
		if bombs == 0 {
			board[x][y] = 'B'

			for i := 0; i < 8; i++ {
				nextX := x + dx[i]
				nextY := y + dy[i]
				if nextX < 0 || nextY < 0 || nextX > m-1 || nextY > n-1 {
					continue
				}

				if had[nextX*n+nextY] {
					continue
				}

				dfs(nextX, nextY)
			}
			return
		}

		board[x][y] = '0' + byte(bombs)
	}

	dfs(x, y)

	return board
}
