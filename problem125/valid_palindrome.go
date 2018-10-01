package problem125

import (
	"strings"
)

func isPalindrome(s string) bool {
	sLen := len(s)
	if sLen == 0 {
		return true
	}

	i, j := 0, sLen-1
	for i < j {

		// 不是字母数字的 跳过
		for i < j && !isValid(s[i]) {
			i++
		}

		for i < j && !isValid(s[j]) {
			j--
		}

		if s[i] != s[j] {
			if s[i] > s[j] && s[j] >= 65 {
				if s[i]-32 == s[j] {
					i++
					j--
					continue
				}
			}

			if s[j] > s[i] && s[i] >= 65 {
				if s[j]-32 == s[i] {
					i++
					j--
					continue
				}
			}

			return false
		}

		i++
		j--
	}

	return true
}

// 字母或数字
func isValid(b byte) bool {
	// 48-57 65-90 97-122
	if b < 48 ||
		(57 < b && b < 65) ||
		(90 < b && b < 97) ||
		122 < b {
		return false
	}

	return true
}

// 另一解法
func isPalindrome2(s string) bool {
	s = strings.ToLower(s)

	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isChar(s[i]) {
			i++
		}

		for i < j && !isChar(s[j]) {
			j--
		}

		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

// 是否合法字符
func isChar(b byte) bool {
	if ('a' <= b && b <= 'z') ||
		('0' <= b && b <= '9') {
		return true
	}

	return false
}
