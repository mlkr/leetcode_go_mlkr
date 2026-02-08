package problem668

func findKthNumber(m int, n int, k int) int {
	// 小的数在前面，节省计算时间
	if m > n {
		m, n = n, m
	}

	lo, hi := 1, m*n
	for lo < hi {
		mid := (lo + hi) / 2
		c := count(mid, m, n)

		if c >= k {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}

// 计算小 <=v 的数有多少个
func count(v, m, n int) int {
	res := 0
	for i := 1; i <= m; i++ {
		res += min(v/i, n)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// 1234
// 123

// 1 2 3 4
// 2 4 6 8
// 3 6 9 12
