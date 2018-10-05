package problem126

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordList = deleteWord(beginWord, wordList)

	// 记录转换关系 n->w
	trans := make(map[string][]string)
	isEnd := false
	dep := 1
	var bfs func(nodes, words []string)
	bfs = func(nodes, words []string) {
		dep++

		newNodes := make([]string, 0, len(wordList))
		newWords := make([]string, 0, len(wordList))
		for _, w := range words {
			canTrans := false

			for _, n := range nodes {
				if isTransable(n, w) {
					canTrans = true
					trans[n] = append(trans[n], w)
				}
			}

			if canTrans {
				newNodes = append(newNodes, w)
				if w == endWord {
					isEnd = true
				}
			} else {
				newWords = append(newWords, w)
			}
		}

		if isEnd ||
			len(newNodes) == 0 ||
			len(newWords) == 0 {
			return
		}

		bfs(newNodes, newWords)
	}

	nodes := []string{beginWord}
	bfs(nodes, wordList)

	res := [][]string{}
	if isEnd == false {
		return res
	}

	path := make([]string, dep)
	path[0] = beginWord
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == dep {
			if path[idx-1] == endWord {
				pathC := make([]string, len(path))
				copy(pathC, path)
				res = append(res, pathC)
			}

			return
		}

		for _, v := range trans[path[idx-1]] {
			path[idx] = v
			dfs(idx + 1)
		}
	}

	dfs(1)

	return res
}

func deleteWord(word string, wordList []string) []string {
	i := 0
	for ; i < len(wordList); i++ {
		if word == wordList[i] {
			break
		}
	}

	if i == len(wordList) {
		return wordList
	}

	wordList[i] = wordList[len(wordList)-1]
	wordList = wordList[:len(wordList)-1]

	return wordList
}

func isTransable(node, word string) bool {
	// 题目给出没有相同的字符串
	once := false
	for i := 0; i < len(node); i++ {
		if node[i] != word[i] {
			if once {
				return false
			}
			once = true
		}
	}

	return true
}
