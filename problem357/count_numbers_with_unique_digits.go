package problem357

// 参看 https://blog.csdn.net/qq508618087/article/details/51656771
// 9 9 8 7 6 ... 依次减少
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 10
	}

	res := 10
	val := 9
	for i := 2; i <= n; i++ {
		val *= 11 - i
		res += val
	}

	return res
}
