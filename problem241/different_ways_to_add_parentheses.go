package problem241

import "strconv"

func diffWaysToCompute(input string) []int {
	cache := make(map[string][]int)

	var dfs func(string) []int
	dfs = func(s string) []int {
		if c, ok := cache[s]; ok {
			return c
		}

		res := make([]int, 0, len(s))
		for i := range s {
			if s[i] == '+' || s[i] == '-' || s[i] == '*' {
				for _, left := range dfs(s[:i]) {
					for _, right := range dfs(s[i+1:]) {
						res = append(res, operate(left, right, s[i]))
					}
				}
			}
		}

		if len(res) == 0 {
			val, _ := strconv.Atoi(s)
			res = append(res, val)
		}

		cache[s] = res

		return res
	}

	return dfs(input)
}

func operate(left, right int, opt byte) int {
	switch opt {
	case '+':
		return left + right
	case '-':
		return left - right
	default:
		return left * right
	}
}
