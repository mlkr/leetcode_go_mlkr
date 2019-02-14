package problem39

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	solution := []int{}
	res := [][]int{}
	cs(candidates, target, solution, &res)

	return res
}

func cs(candidates []int, target int, solution []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, solution)
	}

	if len(candidates) == 0 || target < candidates[0] {
		return
	}

	// 这样处理一下的用意是，让切片的容量等于长度，以后append的时候，会分配新的底层数组
	// 避免多处同时对底层数组进行修改，产生错误的答案。
	// 可以注释掉以下语句，运行单元测试，查看错误发生。
	// 两个冒号 cap len 都设置为 len(solution)
	solution = solution[:len(solution):len(solution)]

	cs(candidates, target-candidates[0], append(solution, candidates[0]), res)
	cs(candidates[1:], target, solution, res)
}
