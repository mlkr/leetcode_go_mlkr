package problem539

import (
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	size := len(timePoints)
	tp := make([]int, size)
	for i, timePoint := range timePoints {
		tp[i] = analy(timePoint)
	}

	sort.Ints(tp)
	res := 24*60 + tp[0] - tp[size-1]

	for i := 1; i < size; i++ {
		if (tp[i] - tp[i-1]) < res {
			res = tp[i] - tp[i-1]
		}

		if res == 0 {
			return res
		}
	}

	return res
}

func analy(timePoint string) int {
	h, _ := strconv.Atoi(timePoint[:2])
	m, _ := strconv.Atoi(timePoint[3:])

	return h*60 + m
}
