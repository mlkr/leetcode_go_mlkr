package problem367

func isPerfectSquare(num int) bool {
	s := sqrt(num)
	return s*s == num
}

// 牛顿法求平方根
func sqrt(num int) int {
	x := num
	for x*x > num {
		x = (x + num/x) / 2
	}

	return x
}

// 二分法求平方根
func isPerfectSquare2(num int) bool {
	l, r := 0, num
	for l < r {
		mid := (l + r) / 2
		if mid*mid < num {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l*l == num
}
