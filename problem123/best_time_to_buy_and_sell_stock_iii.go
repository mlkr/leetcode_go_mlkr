package problem123

func maxProfit(prices []int) int {
	pricesLen := len(prices)
	if pricesLen <= 1 {
		return 0
	}

	profits := []int{}
	profit := 0
	for i := 1; i < pricesLen; i++ {
		priceDiff := prices[i] - prices[i-1]
		if priceDiff*profit >= 0 {
			profit += priceDiff
			continue
		}

		profits = append(profits, profit)
		profit = priceDiff
	}

	profits = append(profits, profit)

	res := 0
	for i := 0; i < len(profits); i++ {
		sum := max(profits[:i]) + max(profits[i:])
		if sum > res {
			res = sum
		}
	}

	return res
}

func max(nums []int) int {
	sum := 0
	maxSum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum < 0 {
			sum = 0
		}

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

// 另一解法(优化上面的解法)
func maxProfit2(prices []int) int {
	pricesLen := len(prices)
	if pricesLen <= 1 {
		return 0
	}

	profits := []int{}
	profit := 0
	for i := 1; i < pricesLen; i++ {
		priceDiff := prices[i] - prices[i-1]
		if priceDiff*profit >= 0 {
			profit += priceDiff
			continue
		}

		profits = append(profits, profit)
		profit = priceDiff
	}

	profits = append(profits, profit)

	return getMaxProfit(profits)
}

func getMaxProfit(profits []int) int {
	size := len(profits)

	toRight := make([]int, size)
	sum := 0
	maxSum := 0
	for i := 0; i < size; i++ {
		sum += profits[i]
		if sum < 0 {
			sum = 0
		}

		if sum > maxSum {
			maxSum = sum
		}
		toRight[i] = maxSum
	}

	toLeft := make([]int, size)
	sum = 0
	maxSum = 0
	for j := size - 1; j >= 0; j-- {
		sum += profits[j]
		if sum < 0 {
			sum = 0
		}

		if sum > maxSum {
			maxSum = sum
		}
		toLeft[j] = maxSum
	}

	res := toRight[size-1]
	for i := 0; i < size-1; i++ {
		sum := toRight[i] + toLeft[i+1]
		if sum > res {
			res = sum
		}
	}

	return res
}
