package problem52

func totalNQueens(n int) int {
	if n == 0 || n == 2 {
		return 0
	}

	if n == 1 {
		return 1
	}

	res := 0
	cols := make([]bool, n)
	d1, d2 := make([]bool, 2*n), make([]bool, 2*n)
	dfs(0, n, cols, d1, d2, &res)

	return res

}

func dfs(r, n int, cols, d1, d2 []bool, res *int) {
	if r == n {
		*res++
		return
	}

	for c := 0; c < n; c++ {
		d1k := r - c + n
		d2k := r + c

		if !cols[c] && !d1[d1k] && !d2[d2k] {
			cols[c], d1[d1k], d2[d2k] = true, true, true
			dfs(r+1, n, cols, d1, d2, res)
			cols[c], d1[d1k], d2[d2k] = false, false, false
		}
	}
}
