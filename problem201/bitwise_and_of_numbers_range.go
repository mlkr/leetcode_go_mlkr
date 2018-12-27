package problem201

func rangeBitwiseAnd(m int, n int) int {
	var count uint
	for m >= 1 && n >= 1 {
		if m == n {
			return m << count
		}

		m >>= 1
		n >>= 1
		count++
	}

	return 0
}
