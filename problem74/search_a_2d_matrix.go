package problem74

func searchMatrix(matrix [][]int, target int) bool {
	rowNum := len(matrix)
	colNum := len(matrix[0])
	if rowNum == 0 {
		return false
	}

	if colNum == 0 {
		return false
	}

	if target < matrix[0][0] || target > matrix[rowNum-1][colNum-1] {
		return false
	}

	r := 0
	for r < rowNum && target >= matrix[r][0] {
		r++
	}
	r--

	i := 0
	j := colNum - 1
	for i <= j {
		mid := (i + j) / 2
		switch {
		case matrix[r][mid] < target:
			i = mid + 1
		case matrix[r][mid] > target:
			j = mid - 1
		default:
			return true
		}
	}

	return false
}
