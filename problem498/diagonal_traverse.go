package problem498

// 我们可以发现，所有的上行都是移动到 [currentRow - 1, currentCol + 1] 的位置，所有的下行都是移动到 [currentRow + 1, currentCol - 1] 的位置。
// 然后再看什么时候转向，当 currentNode 处于矩形的四条边上而且 next node 的位置 out of boundary 的时候我们就要转向了。
func findDiagonalOrder(mat [][]int) []int {
	m := len(mat)
	if m == 0 {
		return []int{}
	}

	n := len(mat[0])
	mn := m * n
	if mn == 0 {
		return []int{}
	}

	isUpping := true
	next := func(i, j int) (int, int) {
		if isUpping {
			// 上升
			i--
			j++

			if j == n {
				j = n - 1
				i += 2
				isUpping = false
			}

			if i == -1 {
				i = 0
				isUpping = false
			}

			return i, j
		}

		// 下降
		i++
		j--

		if i == m {
			i = m - 1
			j += 2
			isUpping = true
		}

		if j == -1 {
			j = 0
			isUpping = true
		}

		return i, j
	}

	res := make([]int, mn)
	i, j := 0, 0

	for k := 0; k < mn; k++ {
		res[k] = mat[i][j]
		i, j = next(i, j)
	}

	return res
}
