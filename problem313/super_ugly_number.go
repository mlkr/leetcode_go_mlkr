package problem313

func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}

	size := len(primes)
	candidates := make([]int, size)
	copy(candidates, primes)

	pos := make([]int, size)

	res := make([]int, n)
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = min(candidates)
		for j := 0; j < size; j++ {
			if candidates[j] == res[i] {
				pos[j]++
				candidates[j] = res[pos[j]] * primes[j]
			}
		}
	}

	return res[n-1]
}

func min(nums []int) int {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	return min
}
