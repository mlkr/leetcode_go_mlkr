package problem475

import "sort"

// 参看https://blog.csdn.net/fuxuemingzhu/article/details/79254295
func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)

	res := 0
	for _, house := range houses {
		i := findHeater(heaters, house)
		if i == 0 {
			res = max(res, abs(heaters[i], house))
			continue
		}

		res = max(res, min(abs(house, heaters[i-1]), abs(heaters[i], house)))
	}

	return res
}

func findHeater(heaters []int, k int) int {
	l := 0
	r := len(heaters) - 1

	for l < r {
		m := (l + r) / 2
		if heaters[m] < k {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
