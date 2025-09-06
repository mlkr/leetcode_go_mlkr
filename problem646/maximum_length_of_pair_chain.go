package problem646

import "sort"

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	num := -1 << 31
	res := 0
	for _, pair := range pairs {
		if num < pair[0] {
			res++
			num = pair[1]
		}
	}

	return res
}
