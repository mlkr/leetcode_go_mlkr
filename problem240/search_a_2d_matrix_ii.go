package problem240

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}

	cols := len(matrix[0])
	if cols == 0 {
		return false
	}

	for r := 0; r < rows; r++ {
		if matrix[r][0] > target {
			break
		}

		if matrix[r][0] <= target && target <= matrix[r][cols-1] {
			if binarySearch(matrix[r], target) {
				return true
			}
		}
	}

	return false
}

func binarySearch(nums []int, target int) bool {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		switch {
		case target < nums[mid]:
			j = mid - 1
		case target > nums[mid]:
			i = mid + 1
		default:
			return true
		}
	}

	return false
}

// 另一解法
func searchMatrix2(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}

	cols := len(matrix[0])
	if cols == 0 {
		return false
	}

	i, j := rows-1, 0
	for i >= 0 && j < cols {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] < target {
			j++
		} else {
			i--
		}
	}

	return false
}
