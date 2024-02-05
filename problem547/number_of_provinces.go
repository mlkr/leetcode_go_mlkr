package problem547

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)

	// 每个城市单独算一个省的话，最多有res个
	res := n

	// 初始每个城市省会都是它自己
	province := make([]int, n)
	for i := range province {
		province[i] = i
	}

	// 将a,b 结合为一个省
	union := func(a, b int) {
		for i := range province {
			if province[i] == a {
				province[i] = b
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				if province[i] != province[j] {
					union(province[i], province[j])
					res--
				}
			}
		}
	}

	return res
}

// 这个更好
func findCircleNum2(isConnected [][]int) int {
	n := len(isConnected)
	visited := make(map[int]bool)

	var dfs func(int)
	dfs = func(i int) {
		visited[i] = true

		for j, v := range isConnected[i] {
			if v == 1 && !visited[j] {
				dfs(j)
			}
		}
	}

	province := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)
			province++
		}
	}

	return province
}
