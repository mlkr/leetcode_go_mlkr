package problem410

func splitArray(nums []int, m int) int {
	l := 0
	r := 0
	for _, num := range nums {
		l = max(l, num)
		r += num
	}

	isBigger := func(guess int) bool {
		sum := 0
		count := 1
		for _, num := range nums {
			sum += num

			if sum > guess {
				count++
				if count > m {
					return false
				}

				sum = num
			}
		}

		return true
	}

	// 二分法
	for l <= r {
		mid := (l + r) / 2
		if isBigger(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
