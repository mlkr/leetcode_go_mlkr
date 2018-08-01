package problem79

func exist(board [][]byte, word string) bool {
	rowCount := len(board)
	if rowCount == 0 {
		return false
	}

	colCount := len(board[0])
	if colCount == 0 {
		return false
	}

	wordLen := len(word)
	if wordLen == 0 {
		return false
	}

	var dfs func(r, c, index int) bool
	dfs = func(r, c, index int) bool {
		if index == wordLen {
			return true
		}

		if r >= rowCount || r < 0 || c >= colCount || c < 0 || board[r][c] != word[index] {
			return false
		}

		temp := board[r][c]
		board[r][c] = 0

		if dfs(r-1, c, index+1) ||
			dfs(r+1, c, index+1) ||
			dfs(r, c-1, index+1) ||
			dfs(r, c+1, index+1) {
			return true
		}

		board[r][c] = temp

		return false
	}

	for r := 0; r < rowCount; r++ {
		for c := 0; c < colCount; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}

	return false
}
