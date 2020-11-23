package problem459

import "strings"

func repeatedSubstringPattern(s string) bool {
	size := len(s)
	if size < 2 {
		return false
	}

	ss := (s + s)[1 : size*2-1]

	return strings.Contains(ss, s)
}

// 最佳
func repeatedSubstringPattern2(s string) bool {
	size := len(s)

	for i := 1; i <= size/2; i++ {
		if size%i != 0 {
			continue
		}

		j := 0
		for j+i < size && s[j] == s[j+i] {
			j++
		}

		if j+i == size {
			return true
		}
	}

	return false
}
