package problem673

func findNumberOfLIS(nums []int) int {
	n := len(nums)

	// lens[i] 以nums[i]结尾的最长递增子序列的长度
	lens := make([]int, n)

	//counts[i] 以nums[i]结尾的最长递增子序列的个数
	counts := make([]int, n)

	maxLen := 0
	res := 0
	for j := 0; j < n; j++ {
		lens[j] = 1
		counts[j] = 1

		for i := 0; i < j; i++ {
			// 递增
			if nums[i] < nums[j] {

				// 证明之前已经有相同长度的递增序列
				// 个数要累加
				if lens[i]+1 == lens[j] {
					counts[j] += counts[i]
				}

				// 更新lens[j]
				if lens[i]+1 > lens[j] {
					lens[j] = lens[i] + 1
					counts[j] = counts[i]
				}
			}
		}

		if maxLen == lens[j] {
			res += counts[j]
		}

		if maxLen < lens[j] {
			maxLen = lens[j]
			res = counts[j]
		}
	}

	return res
}
