package problem338

func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		// 例
		// 5: 101
		// 10: 1010
		// 10 二进制1的个数 = 10&1 + 5 二进制1的个数
		// 11: 1011
		// 11 二进制1的个数 = 11&1 + 5 二进制1的个数
		res[i] = i&1 + res[i>>1]
	}

	return res
}
