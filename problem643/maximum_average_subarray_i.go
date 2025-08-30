package problem643

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	res := sum
	for i := k; i < len(nums); i++ {
		sum -= nums[i-k]
		sum += nums[i]
		res = max(res, sum)
	}

	return float64(res) / float64(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
