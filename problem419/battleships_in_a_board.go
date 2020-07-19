package problem419

func countBattleships(board [][]byte) int {
	rows := len(board)
	if rows == 0 {
		return 0
	}

	cols := len(board[0])
	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'X' && (r == 0 || board[r-1][c] == '.') &&
				(c == 0 || board[r][c-1] == '.') {

				count++
			}
		}
	}

	return count
}
