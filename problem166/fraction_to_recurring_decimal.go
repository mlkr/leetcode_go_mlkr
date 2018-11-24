package problem166

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(n int, d int) string {
	if n == 0 {
		return "0"
	}

	// 判断正负
	if n*d < 0 {
		return "-" + fractionToDecimal(abs(n), abs(d))
	}

	n = abs(n)
	d = abs(d)

	if n > d {
		nextRes := fractionToDecimal(n%d, d)
		return strconv.Itoa(n/d) + nextRes[1:]
	}

	// 记录循环的位置
	loopMap := make(map[int]int, 1024)

	res := make([]byte, 1024)
	res[0] = '0'
	res[1] = '.'

	i := 2
	for n > 0 {

		n *= 10
		index, ok := loopMap[n]
		if ok {
			return fmt.Sprintf("%s(%s)", res[:index], res[index:i])
		}
		loopMap[n] = i

		num := n / d
		n -= (num * d)
		res[i] = byte(num) + '0'

		i++
	}

	return string(res[:i])
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
