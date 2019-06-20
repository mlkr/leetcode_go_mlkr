package problem301

func removeInvalidParentheses(s string) []string {
	var dfs func(first, last int, pair []byte, str string, res []string) []string
	dfs = func(first, last int, pair []byte, str string, res []string) []string {
		count := 0
		for j := last; j < len(str); j++ {
			if str[j] == pair[0] {
				count++
			}

			if str[j] == pair[1] {
				count--
			}

			if count >= 0 {
				continue
			}

			for i := first; i <= j; i++ {
				if str[i] == pair[1] && (i == first || str[i-1] != pair[1]) {
					newStr := str[:i] + str[i+1:]
					res = dfs(i, j, pair, newStr, res)
				}
			}

			return res
		}

		// 反过来再来一次
		newStr := reverse(str)
		if string(pair) == "()" {
			return dfs(0, 0, []byte(")("), newStr, res)
		}

		res = append(res, newStr)
		return res
	}

	res := []string{}
	return dfs(0, 0, []byte("()"), s, res)
}

func reverse(s string) string {
	str := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}

	return string(str)
}
