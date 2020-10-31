package problem455

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	i, j, count := 0, 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			count++
			i++
		}

		j++
	}

	return count
}
