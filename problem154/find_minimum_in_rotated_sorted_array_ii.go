package problem154

func findMin(nums []int) int {
	min := nums[0]
	size := len(nums)
	for i := 1; i < size; i++ {
		if nums[i] < nums[i-1] {
			min = nums[i]
			break
		}
	}

	return min
}
