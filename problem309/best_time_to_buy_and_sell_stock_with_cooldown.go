package problem309

// https://blog.csdn.net/zjuPeco/article/details/76468185
func maxProfit(prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}

	s0 := 0
	s1 := -prices[0]
	s2 := 0
	for i := 1; i < size; i++ {
		pre0, pre1, pre2 := s0, s1, s2
		s0 = max(pre0, pre2)
		s1 = max(pre1, pre0-prices[i])
		s2 = pre1 + prices[i]
	}

	return max(s0, s2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
