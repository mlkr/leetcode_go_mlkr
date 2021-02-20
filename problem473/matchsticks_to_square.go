package problem473

import "sort"

func makesquare(nums []int) bool {
	size := len(nums)
	if size < 4 {
		return false
	}

	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	if sum%4 != 0 {
		return false
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	avg := sum / 4
	if nums[0] > avg {
		return false
	}

	var dfs func(index int, edges [4]int) bool
	dfs = func(index int, edges [4]int) bool {
		if index == size {
			return true
		}

		seen := make(map[int]struct{})
		for i := range edges {
			if edges[i]+nums[index] > avg {
				continue
			}

			if _, ok := seen[edges[i]]; ok {
				continue
			}
			seen[edges[i]] = struct{}{}

			edges[i] += nums[index]
			if dfs(index+1, edges) {
				return true
			}
			edges[i] -= nums[index]
		}

		return false
	}

	return dfs(0, [4]int{})
}
