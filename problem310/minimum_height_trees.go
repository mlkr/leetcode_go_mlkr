package problem310

func findMinHeightTrees(n int, edges [][]int) []int {
	if n < 2 {
		return []int{0}
	}

	link := make([][]int, n)
	for i := 0; i < n; i++ {
		link[i] = make([]int, 0, 5)
	}

	count := make([]int, n)
	for _, edge := range edges {
		link[edge[0]] = append(link[edge[0]], edge[1])
		link[edge[1]] = append(link[edge[1]], edge[0])
		count[edge[0]]++
		count[edge[1]]++
	}

	leafs := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if count[i] == 1 {
			leafs = append(leafs, i)
		}
	}

	// 剪枝
	var leaf, node int
	for n > 2 {
		newLeafs := make([]int, 0, len(leafs))
		for _, leaf = range leafs {
			n--
			for _, node = range link[leaf] {
				count[node]--
				if count[node] == 1 {
					newLeafs = append(newLeafs, node)
				}
			}
		}
		leafs = newLeafs
	}

	return leafs
}
