package problem121

func maxProfit(prices []int) int {
	length := len(prices)
	profit := 0
	if length <= 1 {
		return profit
	}

loop:
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			priceDiff := prices[j] - prices[i]

			if priceDiff < 0 {
				break
			}

			if priceDiff > profit {
				profit = priceDiff
			}

			if j == length-1 {
				break loop
			}
		}
	}

	return profit
}

// 另一解法(简化上面的解法 最佳)
func maxProfit2(prices []int) int {
	length := len(prices)
	profit := 0
	if length <= 1 {
		return profit
	}

	priceDiff := 0
	for i := 1; i < length; i++ {
		priceDiff += prices[i] - prices[i-1]
		if priceDiff < 0 {
			priceDiff = 0
		}

		if priceDiff > profit {
			profit = priceDiff
		}
	}

	return profit
}
