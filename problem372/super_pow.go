package problem372

// (a*b) % c = ((a%c) * (b%c)) % c
// a^1234567 % k = (a^1234560 % k) * (a^7 % k) % k = (a^123456 % k)^10 % k * (a^7 % k) % k
// 详解
// https://leetcode-cn.com/problems/super-pow/solution/you-qian-ru-shen-kuai-su-mi-suan-fa-xiang-jie-by-l/
func superPow(a int, b []int) int {
	base := 1337

	var powmod func(a, k int) int
	powmod = func(a, k int) int {
		if k == 0 {
			return 1
		}

		// k 为奇数
		if k%2 == 1 {
			return (a * powmod(a, k-1)) % base
		}

		// k 为偶数
		sub := powmod(a, k/2)
		return (sub * sub) % base
	}

	size := len(b)
	if size == 0 {
		return 1
	}

	num := b[size-1]
	newB := b[:size-1]

	return (powmod(superPow(a, newB), 10) * powmod(a, num)) % base
}
