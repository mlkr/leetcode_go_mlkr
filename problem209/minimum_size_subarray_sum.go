package problem209

func minSubArrayLen(s int, nums []int) int {
	count := 0
	start := 0
	size := len(nums) + 1
	for i, n := range nums {
		count += n

		for count >= s {
			if (i-start)+1 < size {
				size = (i - start) + 1
			}

			count -= nums[start]
			start++
		}
	}

	if size == len(nums)+1 {
		return 0
	}

	return size
}
