package problem397

func integerReplacement(n int) int {
	rec := make(map[int]int)
	rec[1] = 0

	var dfs func(i int) int
	dfs = func(i int) int {
		if v, ok := rec[i]; ok {
			return v
		}

		if i%2 == 0 {
			rec[i] = dfs(i/2) + 1
			return rec[i]
		}

		return min(dfs(i+1), dfs(i-1)) + 1
	}

	return dfs(n)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}