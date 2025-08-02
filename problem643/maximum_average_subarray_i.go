package problem643

func findMaxAverage(nums []int, k int) float64 {
	count := 0
	sum := 0
	maxSum := -1<<31 + 1
	for i, n := range nums {
		count++
		sum += n
		if k == count {
			maxSum = max(maxSum, sum)
			count--
			sum -= nums[i-k+1]
		}
	}

	return float64(maxSum) / float64(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
