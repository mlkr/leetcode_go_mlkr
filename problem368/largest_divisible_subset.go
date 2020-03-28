package problem368

import "sort"

func largestDivisibleSubset(nums []int) []int {
	size := len(nums)
	if size == 0 {
		return nums
	}

	// 升序排序
	sort.Ints(nums)

	// 保留之前的索引
	pre := make([]int, size)

	// 记录符合要求的子集长度
	dp := make([]int, size)
	for i := range dp {
		dp[i] = 1
	}

	idx, max := 0, 1

	for i := 1; i < size; i++ {
		for j := 0; i > j; j++ {
			if nums[j]<<1 > nums[i] {
				break
			}

			if nums[i]%nums[j] != 0 {
				continue
			}

			if dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				pre[i] = j
			}

			if max < dp[i] {
				max = dp[i]
				idx = i
			}
		}
	}

	res := make([]int, max)
	for i := max - 1; i >= 0; i-- {
		res[i] = nums[idx]
		idx = pre[idx]
	}

	return res
}
