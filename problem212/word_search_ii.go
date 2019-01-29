package problem212

func findWords(board [][]byte, words []string) []string {
	rows := len(board)
	cols := len(board[0])

	res := []string{}
	had := make(map[string]bool, len(words))

	var dfs func(r, c int, words []string, index int)
	dfs = func(r, c int, words []string, index int) {
		if r >= rows || r < 0 || c >= cols || c < 0 {
			return
		}

		newWords := []string{}
		for _, word := range words {
			if word[index] == board[r][c] {
				if len(word)-1 == index {
					if !had[word] {
						had[word] = true
						res = append(res, word)
					}
					continue
				} else {
					newWords = append(newWords, word)
				}
			}
		}

		if len(newWords) == 0 {
			return
		}

		temp := board[r][c]
		board[r][c] = 0

		dfs(r-1, c, newWords, index+1)
		dfs(r+1, c, newWords, index+1)
		dfs(r, c-1, newWords, index+1)
		dfs(r, c+1, newWords, index+1)

		board[r][c] = temp
		return
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dfs(i, j, words, 0)
		}
	}

	return res
}

// 解法二 (最佳)
func findWords2(board [][]byte, words []string) []string {
	rows := len(board)
	cols := len(board[0])

	root := buildTrie(words)
	res := []string{}

	var dfs func(r, c int, trie *Trie)
	dfs = func(r, c int, trie *Trie) {
		b := board[r][c]
		idx := b - 'a'
		if b == '#' || trie.sons[idx] == nil {
			return
		}

		trie = trie.sons[idx]

		if trie.word != "" {
			res = append(res, trie.word)
			trie.word = ""
		}

		board[r][c] = '#'
		if r > 0 {
			dfs(r-1, c, trie)
		}

		if r < rows-1 {
			dfs(r+1, c, trie)
		}

		if c > 0 {
			dfs(r, c-1, trie)

		}

		if c < cols-1 {
			dfs(r, c+1, trie)

		}
		board[r][c] = b

	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dfs(r, c, root)
		}
	}

	return res
}

type Trie struct {
	sons [26]*Trie
	word string
}

func buildTrie(words []string) *Trie {
	root := &Trie{}

	for _, word := range words {
		node := root

		for _, b := range word {
			idx := b - 'a'
			if node.sons[idx] == nil {
				node.sons[idx] = &Trie{}
			}

			node = node.sons[idx]
		}

		node.word = word
	}

	return root
}
