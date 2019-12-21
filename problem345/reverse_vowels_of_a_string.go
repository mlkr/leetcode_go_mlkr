package problem345

func reverseVowels(s string) string {
	b := []byte(s)
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isVowel(b[l]) {
			l++
		}

		for l < r && !isVowel(b[r]) {
			r--
		}

		if l < r {
			b[l], b[r] = b[r], b[l]
		}

		l++
		r--
	}

	return string(b)
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}

	return false
}
