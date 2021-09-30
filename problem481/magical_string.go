package problem481

import "strings"

func magicalString(n int) int {
	bytes := make([]byte, n+2)
	copy(bytes, "122")

	i, t := 2, 2
	for i < n-1 {
		// 下一个填充的是 1 还是 2
		next := (bytes[i] - '0') ^ 3 + '0'

		// 填充的个数
		times := bytes[t] - '0'
		t++

		for times > 0 {
			i++
			bytes[i] = next

			times--
		}
	}

	return strings.Count(string(bytes[:n]), "1")
}
