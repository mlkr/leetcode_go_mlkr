package problem523

func checkSubarraySum(nums []int, k int) bool {
	size := len(nums)
	if size < 2 {
		return false
	}

	sum := 0
	m := make(map[int]int, size)
	m[0] = -1
	for i := range nums {
		sum = (sum + nums[i]) % k

		// 结果与之前的重复，证明已走过一轮，为 k 的倍数
		// 0 始终视为 k 的一个倍数
		if index, ok := m[sum]; ok {
			// 子数组大小 至少为 2
			if i-index > 1 {
				return true
			}
			continue
		}

		m[sum] = i
	}

	return false
}
