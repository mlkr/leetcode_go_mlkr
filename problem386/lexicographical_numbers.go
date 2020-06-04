package problem386

func lexicalOrder(n int) []int {
	res := make([]int, 0, n)

	var dfs func(x int)
	dfs = func(x int) {
		limit := (x + 10) / 10 * 10
		for x < limit && x <= n {
			res = append(res, x)

			if x*10 <= n {
				dfs(x * 10)
			}
			x++
		}
	}

	dfs(1)
	return res
}
