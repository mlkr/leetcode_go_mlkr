package problem68

func fullJustify(words []string, maxWidth int) []string {
	var lineWords []string
	var lineWordsWidth int
	var res []string

	isEnd := false
	for !isEnd {
		words, lineWords, isEnd, lineWordsWidth = split(words, maxWidth)
		res = append(res, combine(lineWords, makeSpace(len(lineWords)-1, maxWidth-lineWordsWidth, isEnd)))
	}

	return res
}

func split(words []string, maxWidth int) ([]string, []string, bool, int) {
	lineWords := make([]string, 1)
	lineWords[0] = words[0]
	width := len(lineWords[0])
	wordCount := len(words)

	i := 1
	for ; i < wordCount; i++ {
		if width+len(lineWords)+len(words[i]) > maxWidth {
			break
		}

		lineWords = append(lineWords, words[i])
		width += len(words[i])
	}

	return words[i:], lineWords, i == wordCount, width
}

func makeSpace(gaps, count int, isEnd bool) []string {
	var res []string

	if isEnd {
		wordCount := gaps + 1
		res = make([]string, wordCount)

		for i := 0; i < gaps; i++ {
			res[i] = " "
		}

		for i := gaps; i < count; i++ {
			res[wordCount-1] += " "

		}

		return res
	}

	if gaps == 0 {
		gaps = 1
	}
	res = make([]string, gaps)

	for i := 0; i < count; i++ {
		res[i%gaps] += " "
	}

	return res
}

func combine(lineWords, space []string) string {
	var res string

	i := 0
	spaceLen := len(space)
	for ; i < spaceLen; i++ {
		res += lineWords[i] + space[i]
	}

	if i < len(lineWords) {
		res += lineWords[i]
	}

	return res
}
