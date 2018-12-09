package problem188

func maxProfit(k int, prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}

	if size/2 <= k {
		return getProfit(prices)
	}

	local := make([]int, k+1)
	global := make([]int, k+1)

	for i := 1; i < size; i++ {
		diff := prices[i] - prices[i-1]
		for j := k; j >= 1; j-- {
			local[j] = max(global[j-1]+max(diff, 0), local[j]+diff)
			global[j] = max(local[j], global[j])
		}
	}

	return global[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func getProfit(prices []int) int {
	size := len(prices)
	profit := 0
	for i := 1; i < size; i++ {
		priceDiff := prices[i] - prices[i-1]
		if priceDiff < 0 {
			priceDiff = 0
		}

		profit += priceDiff
	}

	return profit
}
