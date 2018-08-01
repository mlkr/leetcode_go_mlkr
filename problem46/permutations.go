package problem46

func permute(nums []int) [][]int {
	n := len(nums)

	// 已经使用的数字
	token := make([]bool, n)

	// 一种可能的排序
	seq := make([]int, n)

	// 结果
	ans := make([][]int, 0)

	makePermute(0, n, nums, seq, token, &ans)

	return ans
}

func makePermute(cur int, n int, nums []int, seq []int, token []bool, ans *[][]int) {
	if cur == n {
		tmp := make([]int, n)
		copy(tmp, seq)
		*ans = append(*ans, tmp)
	}

	for i := 0; i < n; i++ {
		if token[i] == false {
			seq[cur] = nums[i]

			token[i] = true
			makePermute(cur+1, n, nums, seq, token, ans)
			token[i] = false
		}
	}
}
