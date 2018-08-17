package problem90

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	var dfs func(idx int, temp []int)
	dfs = func(idx int, temp []int) {
		c := make([]int, len(temp))
		copy(c, temp)
		res = append(res, c)

		for i := idx; i < len(nums); i++ {
			if i == idx || nums[i] != nums[i-1] {
				dfs(i+1, append(temp, nums[i]))
			}
		}
	}

	temp := make([]int, 0, len(nums))
	dfs(0, temp)

	return res
}

func subsetsWithDup2(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	temp := make([]int, 0, len(nums))
	dfs(0, nums, temp, &res)

	return res
}

func dfs(idx int, nums []int, temp []int, res *[][]int) {
	c := make([]int, len(temp))
	copy(c, temp)
	*res = append(*res, c)

	for i := idx; i < len(nums); i++ {
		if i == idx || nums[i] != nums[i-1] {
			dfs(i+1, nums, append(temp, nums[i]), res)
		}
	}
}
