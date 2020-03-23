package problem365

// 贝祖定理
// z = m * x + n * y
// 参看 https://blog.csdn.net/fuxuemingzhu/article/details/83574784
func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}

	if x == 0 || y == 0 {
		return z == 0
	}

	if x+y == z || z == 0 || x == z || y == z {
		return true
	}

	return z%gcd(x, y) == 0
}

// 求最大公约数
func gcd(x, y int) int {

	// y 为较小的数
	if x < y {
		x, y = y, x
	}

	res := x % y
	if res == 0 {
		return y
	}

	return gcd(y, res)
}
