package problem477

func totalHammingDistance(nums []int) int {
	res := 0
	size := len(nums)

	for j := 0; j < 32; j++ {

		ones := 0
		for i := 0; i < size; i++ {
			ones += (nums[i] >> j) & 1
		}

		res += ones * (size - ones)
	}

	return res
}
