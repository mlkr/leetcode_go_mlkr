package problem500

import "strings"

var row1 = map[byte]bool{
	'q': true,
	'w': true,
	'e': true,
	'r': true,
	't': true,
	'y': true,
	'u': true,
	'i': true,
	'o': true,
	'p': true,
}

var row2 = map[byte]bool{
	'a': true,
	's': true,
	'd': true,
	'f': true,
	'g': true,
	'h': true,
	'j': true,
	'k': true,
	'l': true,
}

var row3 = map[byte]bool{
	'z': true,
	'x': true,
	'c': true,
	'v': true,
	'b': true,
	'n': true,
	'm': true,
}

func findWords(words []string) []string {
	res := []string{}
	for _, word := range words {
		lower := strings.ToLower(word)
		if isIn(lower) {
			res = append(res, word)
		}
	}

	return res
}

func isIn(word string) bool {
	b1, b2, b3 := true, true, true
	for i := range word {
		w := word[i]

		if !b1 && !b2 && !b3 {
			break
		}

		if !row1[w] {
			b1 = false
		}

		if !row2[w] {
			b2 = false
		}

		if !row3[w] {
			b3 = false
		}
	}

	return b1 || b2 || b3
}
