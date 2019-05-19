package problem282

func addOperators(num string, target int) []string {
	strs := []string{}

	var dfs func(str, resStr string, res, preNum int)
	dfs = func(str, resStr string, res, preNum int) {
		if len(str) == 0 {
			if res == target {
				strs = append(strs, resStr)
			}
			return
		}

		for i := 1; i <= len(str); i++ {
			curStr := str[:i]
			if curStr[0] == '0' && len(curStr) > 1 {
				// 没有前导 0 的数字
				return
			}

			nextStr := str[i:]
			cur := strToInt(curStr)

			if len(resStr) == 0 {
				dfs(nextStr, curStr, cur, cur)
			} else {
				dfs(nextStr, resStr+"+"+curStr, res+cur, cur)
				dfs(nextStr, resStr+"-"+curStr, res-cur, -cur)
				dfs(nextStr, resStr+"*"+curStr, res-preNum+preNum*cur, preNum*cur)
			}
		}
	}

	dfs(num, "", 0, 0)

	return strs
}

func strToInt(str string) int {
	res := int(str[0] - '0')
	for i := 1; i < len(str); i++ {
		res = res*10 + int(str[i]-'0')
	}

	return res
}
