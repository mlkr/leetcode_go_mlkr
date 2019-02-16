package problem216

func combinationSum3(k int, n int) [][]int {
	used := make(map[int]bool, 10)
	res := [][]int{}
	combination := make([]int, k)

	var dfs func(idx, remain int)
	dfs = func(idx, remain int) {
		if idx == k-1 {
			if remain > combination[idx-1] && remain < 10 && used[remain] == false {
				temp := make([]int, k)
				copy(temp, combination)
				temp[k-1] = remain
				res = append(res, temp)
			}

			return
		}

		min := 0
		if idx != 0 {
			min = combination[idx-1]
		}

		for i := min + 1; i < 10; i++ {
			used[i] = true
			combination[idx] = i
			dfs(idx+1, remain-i)
			used[i] = false
		}

	}

	dfs(0, n)

	return res
}
