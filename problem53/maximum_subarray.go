package problem53

func maxSubArray(nums []int) int {
	len := len(nums)

	if len == 0 {
		return 0
	}

	if len == 1 {
		return nums[0]
	}

	temp, max := nums[0], nums[0]

	i := 1

	for i < len {
		if temp < 0 {
			temp = nums[i]
		} else {
			temp += nums[i]
		}

		if max < temp {
			max = temp
		}

		i++
	}

	return max
}
