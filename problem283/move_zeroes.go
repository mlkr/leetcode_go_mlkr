package problem283

func moveZeroes(nums []int) {
	// i 指向 0, j 指向非0
	i := 0
	size := len(nums)
loop:
	for {
		for nums[i] != 0 {
			i++
			if i >= size {
				break loop
			}
		}

		j := i + 1
		for j < size && nums[j] == 0 {
			j++
		}

		if j >= size {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
}

// 解法二
func moveZeroes2(nums []int) {
	size := len(nums)
	i, j := 0, 0
	for j < size {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}

		j++
	}

	for i < size {
		nums[i] = 0
		i++
	}
}
