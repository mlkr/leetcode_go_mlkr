package problem647

func countSubstrings(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		res += extendPalindrome(s, i, i)
		res += extendPalindrome(s, i, i+1)
	}

	return res
}

// 扩展开 看有多少回文子串
func extendPalindrome(s string, left, right int) int {
	res := 0

	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}

		res++

		left--
		right++
	}

	return res
}
