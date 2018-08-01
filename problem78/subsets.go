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
