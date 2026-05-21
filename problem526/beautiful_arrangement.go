package problem526

func countArrangement(n int) int {
	// 初始化排列
	perm := make([]int, n+1)
	for i := range perm {
		perm[i] = i
	}

	res := 0
	var dfs func(n int)
	dfs = func(n int) {
		// 所有的数都 beautiful
		if n == 0 {
			res++
			return
		}

		for i := n; i > 0; i-- {
			perm[n], perm[i] = perm[i], perm[n]
			if isBeautiful(perm[n], n) {
				dfs(n - 1)
			}
			perm[n], perm[i] = perm[i], perm[n]
		}
	}

	dfs(n)

	return res
}

func isBeautiful(i, j int) bool {
	return i%j == 0 || j%i == 0
}
