package problem674

func findLengthOfLCIS(nums []int) int {
	max, count := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
			if count > max {
				max = count
			}
			continue
		}

		count = 1
	}

	return max
}
