package problem453

// 每次增加除某一个数a之外的所有数，等价与每次减小a
func minMoves(nums []int) int {
	min := nums[0]
	sum := 0
	for i := range nums {
		sum += nums[i]
		if nums[i] < min {
			min = nums[i]
		}
	}

	return sum - len(nums)*min
}
