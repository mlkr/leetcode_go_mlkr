package problem78

func subsets(nums []int) [][]int {
	res := [][]int{}
	temp := []int{}

	recur(nums, temp, &res)

	return res
}

// nums[i...n) 的所有组合情况可以分为两种
//     包含 nums[i] 的
//     不包含 nums[i] 的
// 上面的递推关系，实际上就是 DP 思路。
func recur(nums []int, temp []int, res *[][]int) {
	n := len(nums)
	if n == 0 {
		// sort.Ints(temp)
		*res = append(*res, temp)
		return
	}

	recur(nums[:n-1], temp, res)
	recur(nums[:n-1], append([]int{nums[n-1]}, temp...), res)
}

func subsets2(nums []int) [][]int {
	res := [][]int{}

	var dfs func(int, []int)
	dfs = func(idx int, temp []int) {
		t := make([]int, len(temp))
		copy(t, temp)
		res = append(res, t)

		for i := idx; i < len(nums); i++ {
			dfs(i+1, append(temp, nums[i]))
		}
	}

	temp := make([]int, 0, len(nums))
	dfs(0, temp)

	return res
}
