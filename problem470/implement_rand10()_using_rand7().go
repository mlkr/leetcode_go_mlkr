package problem470

import "math/rand"

// 参看https://www.cnblogs.com/grandyang/p/9727206.html
func rand10() int {
	num := 40
	// 虽然 num范围[0,48], 但是只取[0,39]
	// 因为对于%10运算来说[0,39]取得的结果概率一样， 而且离48最近
	for num > 39 {
		// 这样做是为了每次出现的数概率相等
		// num范围[0,48]
		num = 7*(rand7()-1) + (rand7() - 1)
	}

	return num%10 + 1
}

// 范围[1,7]
func rand7() int {
	return rand.Intn(7) + 1
}

// 0 1 2 3 4 5 6
// 0 7 14 21 28 35 42
// 每次出现的数概率相等, 范围[0,48]
