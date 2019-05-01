package problem264

func nthUglyNumber(n int) int {
	// 1 2 3 4 5 6 都是 ugly number
	if n < 7 {
		return n
	}

	uglys := make([]int, n)
	uglys[0] = 1
	pos := make([]int, 3)
	for i := 1; i < n; i++ {
		a := uglys[pos[0]] * 2
		b := uglys[pos[1]] * 3
		c := uglys[pos[2]] * 5
		min := getMin(a, getMin(b, c))
		if min == a {
			pos[0]++
		}
		if min == b {
			pos[1]++
		}
		if min == c {
			pos[2]++
		}

		uglys[i] = min

	}

	return uglys[n-1]
}

func getMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
