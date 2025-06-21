package problem639

import (
	"strconv"
)

const mod = 1e9 + 7

var m = [144]int{}

func init() {

	setVal := func(s string, val int) {
		key := 0
		for _, r := range s {
			key *= 12
			if r == '*' {
				key += 11
			} else {
				// 把'0'考虑进去，所以 +1
				// 例如 "106"
				// 	如果不 +1，到 "06" 的时候会把 '0' 忽略掉 "06" == "6"
				key += int(r-'0') + 1
			}
		}

		m[key] = val
	}

	for i := 1; i <= 26; i++ {
		setVal(strconv.Itoa(i), 1)
	}

	// 1-9
	setVal("*", 9)

	// 11-19 21-26
	setVal("**", 15)

	setVal("1*", 9)
	setVal("2*", 6)

	for _, r := range "0123456" {
		setVal("*"+string(r), 2)
	}

	for _, r := range "789" {
		setVal("*"+string(r), 1)
	}
}

func numDecodings(s string) int {

	// 所有的排列组合到当前的r位置都会有两种情况（
	// 因为只有26个字母，不会出现三位数代表一个字母的情况）：
	// 	1.r 为一位
	// 	2.r 和前面一位组成两位
	// 例如：11106 到最后一位可以这样分：
	// 	1. 1110 6
	// 	2. 111 06
	// 以此类推倒数第二位可以这样分：
	// 	1. 111 0 6
	// 	2. 11 10 6
	// back1 为 r一位数时 r前面的部分的组合数（可解码数）
	// back2 为 r和前面一位组成两位时 前面的部分的组合数（可解码数）
	// take1 r一位数时 可解码数
	// take2 r和前面一位组成两位时 可解码数
	// 字符串到当前的r位置可解码数为
	// 	cur = back1*take1 + back2*take2

	back1, back2 := 1, 0

	key := 0
	cur := 0
	for _, r := range s {
		key *= 12
		if r == '*' {
			key += 11
		} else {
			key += int(r-'0') + 1
		}

		take2 := m[key]

		key = key % 12
		take1 := m[key]

		cur = (back1*take1 + back2*take2) % mod
		back1, back2 = cur, back1

	}

	return cur
}
