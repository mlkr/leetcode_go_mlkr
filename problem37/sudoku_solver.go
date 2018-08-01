package problem37

func solveSudoku(board [][]byte) {
	if !fill(board, 0, '1') {
		panic("数独无解")
	}
}

func fill(board [][]byte, block int, num byte) bool {
	if block == 9 {
		return true
	}

	if num == '9'+1 {
		return fill(board, block+1, '1')
	}

	rowBegin := (block / 3) * 3
	colBegin := (block % 3) * 3

	// 如果该 num 在 block 中已存在, 填充下一个
	for r := rowBegin; r < rowBegin+3; r++ {
		for c := colBegin; c < colBegin+3; c++ {
			if board[r][c] == num {
				return fill(board, block, num+1)
			}
		}
	}

	// 填充 num
	for r := rowBegin; r < rowBegin+3; r++ {
		for c := colBegin; c < colBegin+3; c++ {
			if checkNum(board, num, r, c) {
				board[r][c] = num
				if fill(board, block, num+1) {
					return true
				}

				board[r][c] = '.'
			}
		}
	}

	return false

}

// 检查数字是否合适
func checkNum(board [][]byte, num byte, row int, col int) bool {

	if board[row][col] != '.' {
		return false
	}

	for i := 0; i < 9; i++ {
		if board[i][col] == num || board[row][i] == num {
			return false
		}
	}

	return true
}
