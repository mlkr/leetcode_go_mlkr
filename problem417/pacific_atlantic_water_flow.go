package problem417

func pacificAtlantic(matrix [][]int) [][]int {
	res := [][]int{}

	rows := len(matrix)
	if rows == 0 {
		return res
	}

	cols := len(matrix[0])
	if cols == 0 {
		return res
	}

	// p 可以到达 pacific 的点
	// a 可以到达 atlantic 的点
	p, a := make([][]bool, rows), make([][]bool, rows)
	for i := 0; i < rows; i++ {
		p[i] = make([]bool, cols)
		a[i] = make([]bool, cols)
	}

	pQue, aQue := [][]int{}, [][]int{}

	for i := 0; i < rows; i++ {
		p[i][0] = true
		a[i][cols-1] = true

		pQue = append(pQue, []int{i, 0})
		aQue = append(aQue, []int{i, cols - 1})
	}

	for i := 0; i < cols; i++ {
		p[0][i] = true
		a[rows-1][i] = true

		pQue = append(pQue, []int{0, i})
		aQue = append(aQue, []int{rows - 1, i})
	}

	dr := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

	bfs := func(queue [][]int, rec [][]bool) {
		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]

			for _, d := range dr {
				r := cell[0] + d[0]
				c := cell[1] + d[1]

				if 0 <= r && r < rows &&
					0 <= c && c < cols &&
					!rec[r][c] &&
					matrix[r][c] >= matrix[cell[0]][cell[1]] {

					rec[r][c] = true
					queue = append(queue, []int{r, c})
				}
			}

		}
	}

	bfs(pQue, p)
	bfs(aQue, a)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if p[r][c] && a[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}

	return res

}
