package problem51

func solveNQueens(n int) [][]string {
	res := [][]string{}

	if n == 0 {
		return res
	}

	// 每列占用
	cols := make([]bool, n)

	// '\' 占用 45度角, 斜率为1
	// y = kx+b (k=1 斜率为1)
	// y = x+b
	// y - x = b
	// 为了保证数组下标 不为负数 y-x+n = b+n
	d1 := make([]bool, 2*n)

	// '/' 占用 135度角, 斜率为-1
	// y = kx+b (k=-1 斜率为-1)
	// y = -x+b
	// y+x = b
	d2 := make([]bool, 2*n)

	board := make([]string, n)

	dfs(0, cols, d1, d2, board, &res, n)

	return res
}

func dfs(r int, cols, d1, d2 []bool, board []string, res *[][]string, n int) {
	if r == n {
		tmp := make([]string, n)
		copy(tmp, board)
		*res = append(*res, tmp)

		return
	}

	for c := 0; c < n; c++ {
		d1k := r - c + n
		// d2k := 2*n - 1 - r - c
		d2k := r + c

		if cols[c] == false && d1[d1k] == false && d2[d2k] == false {
			b := make([]byte, n)
			for i := range b {
				b[i] = '.'
			}

			b[c] = 'Q'
			board[r] = string(b)

			cols[c], d1[d1k], d2[d2k] = true, true, true
			dfs(r+1, cols, d1, d2, board, res, n)
			cols[c], d1[d1k], d2[d2k] = false, false, false

		}
	}
}
