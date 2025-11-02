package problem650

func minSteps(n int) int {
	if n == 1 {
		return 0
	}

	// 2,3 都是素数，素数只能一个一个复制
	if n <= 3 {
		return n
	}

	// n 最小因数是多少？
	// n=12 minSteps=7
	// 111111 111111
	// 111 111
	// 1 1 1
	i := 2
	for i*i <= n {
		if n%i == 0 {
			sub := n / i
			return minSteps(sub) + i
		}

		i++
	}

	// 素数
	return n
}
