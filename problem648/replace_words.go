package problem648

import (
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	hasRoot := make(map[string]bool)
	max, min := len(dictionary[0]), len(dictionary[0])
	for _, root := range dictionary {
		hasRoot[root] = true

		size := len(root)
		if min > size {
			min = size
		}

		if max < size {
			max = size
		}
	}

	words := strings.Fields(sentence)
	for i, word := range words {
		if len(word) < min {
			continue
		}

		ceil := getMin(max, len(word))
		for j := min; j <= ceil; j++ {
			if hasRoot[word[:j]] {
				words[i] = word[:j]
				break
			}
		}
	}

	return strings.Join(words, " ")
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
