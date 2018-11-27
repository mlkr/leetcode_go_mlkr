package problem169

func majorityElement(nums []int) int {
	size := len(nums)
	half := size / 2
	m := make(map[int]int)

	for i := 0; i < size; i++ {
		times := m[nums[i]]
		times++
		if times > half {
			return nums[i]
		}
		m[nums[i]] = times
	}

	return -1
}

// 解法二(最佳)
// 两两不同地抵消掉, 剩下的就是出现次数多于 n/2 的数
func majorityElement2(nums []int) int {
	times, num := 1, nums[0]
	size := len(nums)
	for i := 1; i < size; i++ {
		switch {
		case nums[i] == num:
			times++
		case times > 0:
			times--
		default:
			times, num = 1, nums[i]
		}
	}

	return num
}
