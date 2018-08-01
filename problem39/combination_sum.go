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

	solution = solution[:len(solution):len(solution)]

	cs(candidates, target-candidates[0], append(solution, candidates[0]), res)
	cs(candidates[1:], target, solution, res)
}
