package problem36

func isValidSudoku(board [][]byte) bool {

	// 每行检查
	for x := 0; x < 9; x++ {
		var nums [10]bool
		for y := 0; y < 9; y++ {
			num := covertToNum(board[x][y])
			if num < 0 {
				continue
			}

			if nums[num] { //已经存在
				return false
			}

			nums[num] = true
		}
	}

	// 每列检查
	for y := 0; y < 9; y++ {
		var nums [10]bool
		for x := 0; x < 9; x++ {
			num := covertToNum(board[x][y])
			if num < 0 {
				continue
			}

			if nums[num] { //已经存在
				return false
			}

			nums[num] = true
		}
	}

	// 每个3x3框检查一遍
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var nums [10]bool
			for x := i * 3; x < (i+1)*3; x++ {
				for y := j * 3; y < (j+1)*3; y++ {
					num := covertToNum(board[x][y])
					if num < 0 {
						continue
					}

					if nums[num] { //已经存在
						return false
					}

					nums[num] = true
				}
			}
		}
	}

	return true
}

// 转换数字
func covertToNum(b byte) int {
	if b == '.' {
		return -1
	}

	return int(b - '0')
}
