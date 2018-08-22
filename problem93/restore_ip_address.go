package problem93

import (
	"fmt"
)

func restoreIpAddresses(s string) []string {
	sLen := len(s)
	if sLen < 4 || sLen > 12 {
		return []string{}
	}

	res := make([]string, 4)
	ips := []string{}

	var dfs func(idx, begin int)
	dfs = func(idx, begin int) {
		if idx == 3 {
			if isok(s[begin:sLen]) {
				res[idx] = s[begin:sLen]
				ips = append(ips, ip(res))
			}
			return
		}

		// 最多剩余多少位数字
		remainMax := 3 * (3 - idx)
		for end := begin + 1; end < begin+4; end++ {
			if end+remainMax < sLen {
				continue
			}

			if end > sLen-1 {
				break
			}

			if isok(s[begin:end]) {
				res[idx] = s[begin:end]
				dfs(idx+1, end)
			}

		}
	}

	dfs(0, 0)

	return ips
}

// 检查每个IP的地址段是否正确
func isok(s string) bool {
	sLen := len(s)

	// "0" true
	// "01" false
	if sLen > 1 && s[0] == '0' {
		return false
	}

	if sLen < 3 {
		return true
	}

	// sLen == 3
	switch {
	case s[0] == '1':
		return true
	case s[0] == '2' && s[1] <= '4':
		return true
	case s[0] == '2' && s[1] == '5':
		if s[2] <= '5' {
			return true
		}

		return false
	}

	return false
}

// 输出格式化后的ip
func ip(s []string) string {
	return fmt.Sprintf("%s.%s.%s.%s", s[0], s[1], s[2], s[3])
}
