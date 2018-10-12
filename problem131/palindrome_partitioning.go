package problem131

func partition(s string) [][]string {
	cur := make([]string, 0, len(s))
	res := make([][]string, 0)
	dfs(s, 0, cur, &res)

	return res
}

func dfs(s string, i int, cur []string, res *[][]string) {
	if i == len(s) {
		temp := make([]string, len(cur))
		copy(temp, cur)
		*res = append(*res, temp)
		return
	}

	for j := i; j < len(s); j++ {
		if isPalindrome(s[i : j+1]) {
			dfs(s, j+1, append(cur, s[i:j+1]), res)
		}
	}
}

// 验证回文字符串
func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}
