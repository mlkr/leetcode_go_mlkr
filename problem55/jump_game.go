package problem55

func canJump(nums []int) bool {
loop1:
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] != 0 {
			continue
		}

		j := i - 1
		for ; j >= 0; j-- {
			if nums[j]+j > i {
				continue loop1
			}
		}

		if j == -1 {
			return false
		}

	}

	return true
}
