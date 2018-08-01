package problem40

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
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

	if len(candidates) == 0 || candidates[0] > target {
		return
	}

	solution = solution[:len(solution):len(solution)]

	cs(candidates[1:], target-candidates[0], append(solution, candidates[0]), res)
	cs(next(candidates), target, solution, res)
}

func next(candidates []int) []int {
	i := 0
	for i < len(candidates)-1 && candidates[i] == candidates[i+1] {
		i++
	}

	return candidates[i+1:]
}
