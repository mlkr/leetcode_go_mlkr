package problem77

func combine(n int, k int) [][]int {
	res := [][]int{}
	comb := make([]int, k)

	var dfs func(int, int)
	dfs = func(idx int, begin int) {
		if idx == k {
			temp := make([]int, k)
			copy(temp, comb)
			res = append(res, temp)
			return
		}

		for ; begin <= n-(k-(idx+1)); begin++ {
			comb[idx] = begin
			dfs(idx+1, begin+1)
		}
	}

	dfs(0, 1)

	return res
}
