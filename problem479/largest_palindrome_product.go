package problem479

import "math"

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}

	base := int(math.Pow10(n)) - 1
	p := 0
	for {
		p = makePali(base)
		if isProduct(p, n) {
			break
		}

		base--
	}

	return p % 1337
}

// 得回文数字
func makePali(base int) int {
	p := base
	for base > 0 {
		p = p*10 + base%10
		base /= 10
	}

	return p
}

// 是否为两个数的乘积
func isProduct(p, n int) bool {
	for i := int(math.Pow10(n) - 1); i >= int(math.Pow10(n-1)) && i >= int(math.Sqrt(float64(p))); i-- {
		if p%i == 0 && getDigits(p/i) == n {
			return true
		}
	}

	return false
}

func getDigits(num int) int {
	res := 0
	for num > 0 {
		num /= 10
		res++
	}

	return res
}
