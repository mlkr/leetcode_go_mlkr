package problem60

func getPermutation(n int, k int) string {
	if n == 0 {
		return ""
	}

	res := make([]byte, n)

	// 待取的字符
	rec := make([]byte, n)
	for i := 0; i < n; i++ {
		rec[i] = byte(i) + '1'
	}

	// 第一个位置的字符确定的话, 有(n-1)!种结果
	base := 1
	for i := 2; i <= n-1; i++ {
		base *= i
	}

	k--

	for i := 0; i < n-1; i++ {
		idx := k / base
		res[i] = rec[idx]

		// 更新rec, 去除已使用的rec[idx]
		rec = append(rec[:idx], rec[idx+1:]...)

		// 更新k
		k %= base

		// 更新 base
		base /= (n - 1 - i)
	}

	// 最后一个字符
	res[n-1] = rec[0]

	return string(res)
}
