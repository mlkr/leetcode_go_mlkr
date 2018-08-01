package problem73

func setZeroes(matrix [][]int) {
	rows := make(map[int]int)
	cols := make(map[int]int)

	n := len(matrix[0])

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rows[i] = i
				cols[j] = j
			}
		}
	}

	for i := range rows {
		matrix[i] = make([]int, n)
	}

	for j := range cols {
		for i := range matrix {
			matrix[i][j] = 0
		}
	}
}
