package problem231

func isPowerOfTwo(n int) bool {
	res := false
	var i uint

	for i = 0; i < 32; i++ {
		num := 1 << i
		if num == n {
			res = true
			break
		}

		if num > n {
			break
		}
	}

	return res
}
