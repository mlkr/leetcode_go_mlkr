package problem152

func maxProduct(nums []int) int {
	cur, neg, max := 1, 1, nums[0]
	for _, num := range nums {
		switch {
		case num > 0:
			cur, neg = cur*num, neg*num
		case num < 0:
			cur, neg = neg*num, cur*num
		default:
			cur, neg = 0, 1
		}

		if cur > max {
			max = cur
		}

		if cur <= 0 {
			cur = 1
		}
	}

	return max
}
