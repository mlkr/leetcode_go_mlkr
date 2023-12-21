package problem542

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])

	MAX := m * n
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				continue
			}

			mat[i][j] = MAX

			if i-1 >= 0 {
				mat[i][j] = min(mat[i][j], mat[i-1][j]+1)
			}

			if j-1 >= 0 {
				mat[i][j] = min(mat[i][j], mat[i][j-1]+1)
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if mat[i][j] == 0 {
				continue
			}

			if i+1 < m {
				mat[i][j] = min(mat[i][j], mat[i+1][j]+1)
			}

			if j+1 < n {
				mat[i][j] = min(mat[i][j], mat[i][j+1]+1)
			}
		}
	}

	return mat
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

type Point struct {
	i, j int
}

func updateMatrix2(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	MAX := m * n

	que := []Point{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				que = append(que, Point{i, j})
			} else {
				mat[i][j] = MAX
			}
		}
	}

	for len(que) > 0 {
		i, j := que[0].i, que[0].j

		if i-1 >= 0 && mat[i][j]+1 < mat[i-1][j] {
			mat[i-1][j] = mat[i][j] + 1
			que = append(que, Point{i - 1, j})
		}

		if i+1 < m && mat[i][j]+1 < mat[i+1][j] {
			mat[i+1][j] = mat[i][j] + 1
			que = append(que, Point{i + 1, j})
		}

		if j-1 >= 0 && mat[i][j]+1 < mat[i][j-1] {
			mat[i][j-1] = mat[i][j] + 1
			que = append(que, Point{i, j - 1})
		}

		if j+1 < n && mat[i][j]+1 < mat[i][j+1] {
			mat[i][j+1] = mat[i][j] + 1
			que = append(que, Point{i, j + 1})
		}

		que = que[1:]
	}

	return mat
}
