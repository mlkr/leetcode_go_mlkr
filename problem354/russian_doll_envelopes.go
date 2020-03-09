package problem354

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			// 宽度相同, 高度从大到小排
			return envelopes[i][1] > envelopes[j][1]
		}

		// 宽度从小到大排
		return envelopes[i][0] < envelopes[j][0]
	})

	dp := []int{}
	for i := range envelopes {
		size := len(dp)
		if size == 0 || dp[size-1] < envelopes[i][1] {
			dp = append(dp, envelopes[i][1])
			continue
		}

		index := binSearch(dp, envelopes[i][1])
		dp[index] = envelopes[i][1]
	}

	return len(dp)
}

// 二分查找
func binSearch(dp []int, elem int) int {
	lo, hi := 0, len(dp)-1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if dp[mid] == elem {
			return mid
		}

		if dp[mid] < elem {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return lo
}
