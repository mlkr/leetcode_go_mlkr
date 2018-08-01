package problem70

func climbStairs(n int) int {
	var res []int
	if n+1 < 3 {
		res = make([]int, 3)
	} else {
		res = make([]int, n+1)

	}

	res[0] = 0
	res[1] = 1
	res[2] = 2

	if n < 3 {
		return res[n]
	}

	for i := 3; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}

	return res[n]
}
