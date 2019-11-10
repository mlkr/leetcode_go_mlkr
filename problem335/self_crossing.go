package problem335

//参看 https://blog.csdn.net/zdavb/article/details/51006487
func isSelfCrossing(x []int) bool {
	size := len(x)
	if size < 4 {
		return false
	}

	for i := 3; i < size; i++ {
		if x[i] >= x[i-2] && x[i-1] <= x[i-3] {
			return true
		}

		if i >= 4 {
			if x[i-1] == x[i-3] && x[i]+x[i-4] >= x[i-2] {
				return true
			}
		}

		if i >= 5 {
			if x[i-2] > x[i-4] && x[i]+x[i-4] >= x[i-2] &&
				x[i-1]+x[i-5] >= x[i-3] && x[i-1] < x[i-3] {
				return true
			}
		}
	}

	return false
}
