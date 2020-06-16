package problem394

import (
	"strings"
)

func decodeString(s string) string {

	var dfs func(idx int) (int, string)
	dfs = func(idx int) (int, string) {
		str := ""
		times := 0

		for ; idx < len(s); idx++ {
			switch {
			case s[idx] >= '0' && s[idx] <= '9':
				times = times*10 + int(s[idx]-'0')
			case s[idx] == '[':
				idx_, s_ := dfs(idx + 1)
				str += strings.Repeat(s_, times)
				times = 0
				idx = idx_
			case s[idx] == ']':
				return idx, str
			default:
				str += string(s[idx])
			}
		}

		return idx, str
	}

	_, res := dfs(0)

	return res
}
