package problem47

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	token := make([]bool, n)
	seq := make([]int, n)
	var ans [][]int

	makePermuteUnique(nums, 0, n, token, seq, &ans)

	return ans
}

func makePermuteUnique(nums []int, cur int, n int, token []bool, seq []int, ans *[][]int) {
	if cur == n {
		temp := make([]int, n)
		copy(temp, seq)
		*ans = append(*ans, temp)
	}

	used := make(map[int]bool, n-cur)

	for i := 0; i < n; i++ {
		if token[i] == false && used[nums[i]] == false {
			seq[cur] = nums[i]
			used[nums[i]] = true

			token[i] = true
			makePermuteUnique(nums, cur+1, n, token, seq, ans)
			token[i] = false

		}
	}
}
