package problem54

func spiralOrder(matrix [][]int) []int {
	r := len(matrix)

	if r == 0 {
		return []int{}
	}

	if r == 1 {
		return matrix[0]
	}

	c := len(matrix[0])

	if c == 0 {
		return []int{}
	}

	// res := make([]int, r*c)
	var res []int

	// →
	res = append(res, matrix[0]...)

	// ↓
	for i := 1; i < r-1; i++ {
		res = append(res, matrix[i][c-1])
	}

	// ←
	for j := c - 1; j >= 0; j-- {
		res = append(res, matrix[r-1][j])
	}

	// ↑
	if c > 1 {
		for i := r - 2; i > 0; i-- {
			res = append(res, matrix[i][0])
		}
	}

	if r == 2 || c <= 2 {
		return res
	}

	// 中间部分
	// nextMatrix := make([][]int, r-2)
	var nextMatrix [][]int
	for i := 1; i < r-1; i++ {
		nextMatrix = append(nextMatrix, matrix[i][1:c-1])
	}

	res = append(res, spiralOrder(nextMatrix)...)

	return res

}
