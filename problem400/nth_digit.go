package problem400

// 参看
// https://leetcode-cn.com/problems/nth-digit/solution/zhao-gui-lu-jian-ji-de-jie-jue-ci-wen-ti-by-woxiao/
// [1~9]	 9*1
// [10~99]	 90*2
// [100~999] 900*3
func findNthDigit(n int) int {
	count := 9
	bit := 1
	num := 1
	for n-count*bit > 0 {
		n -= count * bit
		count *= 10
		bit++
		num *= 10
	}

	num = num - 1 + ceil(n, bit)

	index := (n - 1) % bit
	res := num
	for i := bit - 1; i >= index; i-- {
		res = num % 10
		num /= 10
	}

	return res
}

func ceil(a, b int) int {
	num := a / b
	if num*b < a {
		return num + 1
	}

	return num
}
